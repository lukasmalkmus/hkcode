package testdata

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed text_golden.png
	textGolden []byte
	//go:embed qr_golden.png
	qrGolden []byte
	//go:embed qr_boxed_golden.png
	qrBoxedGolden []byte
)

// GetGoldenTextImage returns the golden image for the text based Apple HomeKit®
// setup code.
func GetGoldenTextImage(t *testing.T) image.Image {
	golden, err := png.Decode(bytes.NewReader(textGolden))
	require.NoError(t, err)

	return golden
}

// GetGoldenQRCodeImage returns the golden image for the QR code based Apple
// HomeKit® setup code.
func GetGoldenQRCodeImage(t *testing.T) image.Image {
	golden, err := png.Decode(bytes.NewReader(qrGolden))
	require.NoError(t, err)

	return golden
}

// GetGoldenBoxedQRCodeImage returns the golden image for the boxed QR code
// based Apple HomeKit® setup code.
func GetGoldenBoxedQRCodeImage(t *testing.T) image.Image {
	golden, err := png.Decode(bytes.NewReader(qrBoxedGolden))
	require.NoError(t, err)

	return golden
}
