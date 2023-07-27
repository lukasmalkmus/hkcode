package qr

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/skip2/go-qrcode"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	"github.com/lukasmalkmus/hkcode/hk"
	"github.com/lukasmalkmus/hkcode/internal/assets"
	embeddedFont "github.com/lukasmalkmus/hkcode/internal/assets/font"
)

// CreateCode creates a QR code based Apple HomeKit® setup code.
func CreateCode(setupCode hk.Code, setupID hk.ID, setupFlags hk.Flag, category hk.Category) (image.Image, error) {
	payload, err := CreatePayload(setupCode, setupID, setupFlags, category)
	if err != nil {
		return nil, fmt.Errorf("create payload: %w", err)
	}

	qrc, err := qrcode.New(payload, qrcode.High)
	if err != nil {
		return nil, fmt.Errorf("create QR code: %w", err)
	}

	return qrc.Image(256), nil
}

// CreateBoxedCode creates a QR code based Apple HomeKit® setup code that is
// placed inside a bordered box with the Apple HomeKit® logo and the setup code
// in plain text. These codes are usually found as stickers on MFi accessories.
func CreateBoxedCode(setupCode hk.Code, setupID hk.ID, setupFlags hk.Flag, category hk.Category) (image.Image, error) {
	payload, err := CreatePayload(setupCode, setupID, setupFlags, category)
	if err != nil {
		return nil, fmt.Errorf("create payload: %w", err)
	}

	img, err := assets.Box()
	if err != nil {
		return nil, fmt.Errorf("load box template image: %w", err)
	}

	dimg, ok := img.(draw.Image)
	if !ok {
		return nil, fmt.Errorf("%T is not a drawable image type", img)
	}

	otf, err := embeddedFont.SFMonoBold()
	if err != nil {
		return nil, fmt.Errorf("load font: %w", err)
	}

	qrc, err := qrcode.New(payload, qrcode.Medium)
	if err != nil {
		return nil, fmt.Errorf("create QR code: %w", err)
	}
	qrc.DisableBorder = true
	qrc.BackgroundColor = color.Transparent

	qrImg := qrc.Image(320)
	offset := image.Point{X: 40, Y: 180}

	draw.Draw(dimg, qrImg.Bounds().Add(offset), qrImg, image.Point{}, draw.Src)

	face, err := opentype.NewFace(otf, &opentype.FaceOptions{
		Size:    69,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, fmt.Errorf("create font face: %w", err)
	}
	defer face.Close()

	fd := font.Drawer{
		Dst:  dimg,
		Src:  image.Black,
		Face: face,
	}

	codeStr := setupCode.String()

	for i := 0; i < 4; i++ {
		fd.Dot = fixed.Point26_6{
			X: fixed.I(173 + i*49),
			Y: face.Metrics().Ascent + fixed.I(25),
		}
		fd.DrawString(string(codeStr[i]))

		fd.Dot = fixed.Point26_6{
			X: fixed.I(173 + i*49),
			Y: face.Metrics().Ascent + fixed.I(88),
		}
		fd.DrawString(string(codeStr[i+4]))
	}

	return img, nil
}

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// CreatePayload creates a QR code payload for the given Apple HomeKit® setup
// code.
func CreatePayload(setupCode hk.Code, setupID hk.ID, setupFlags hk.Flag, category hk.Category) (string, error) {
	// Validate and normalize input.
	if !setupCode.Valid() {
		return "", hk.ErrInvalidCode
	} else if !setupID.Valid() {
		return "", hk.ErrInvalidID
	}

	// Changing these will break the code as it will not be recognized by
	// Apple Devices anymore.
	const (
		version  = 0
		reserved = 0
	)

	// Bits 45-43: The "Version" field (0-7). Always set to 0.
	var payload uint64
	payload |= (version & 0x7)

	// Bits 42-39: The "Reserved" field (0-15). Always set to 0.
	payload <<= 4
	payload |= (reserved & 0xf)

	// Bits 38-31: The accessory type (0-255).
	payload <<= 8
	payload |= (uint64(category) & 0xff)

	// Bits 30-27: The setup flags (supported pairing methods, 0-15). Seem to be
	// ignored by Apple devices.
	// Bit 30: Always set to 0.
	// Bit 29: Set to 1 if BLTE pairing is supported, else set to 0.
	// Bit 28: Set to 1 if IP pairing is supported, else set to 0.
	// Bit 27: Set to 1 if NFC pairing is supported, else set to 0.
	payload <<= 4
	payload |= (uint64(setupFlags) & 0xf)

	// Bits 26-0 - The 8-digit setup code (from 0-99999999).
	payload <<= 27
	payload |= (uint64(setupCode) & 0x7fffffff)

	// The result must be 9 digits. If less, pad with leading zeros. Encode as
	// Base 36.
	encodedPayload := make([]byte, 9)
	for i := 0; i < len(encodedPayload); i++ {
		reverseIdx := len(encodedPayload) - i - 1
		encodedPayload[reverseIdx] = base36[payload%36]
		payload /= 36
	}

	return fmt.Sprintf("X-HM://%s%s", encodedPayload, setupID), nil
}
