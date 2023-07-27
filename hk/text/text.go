package text

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	"github.com/lukasmalkmus/hkcode/hk"
	embeddedFont "github.com/lukasmalkmus/hkcode/internal/assets/font"
)

// CreateCode creates a text based Apple HomeKit® setup code.
func CreateCode(setupCode hk.Code) (image.Image, error) {
	if !setupCode.Valid() {
		return nil, hk.ErrInvalidCode
	}

	// 150x50px with 10px padding and white background.
	img := image.NewRGBA(image.Rect(0, 0, 160, 60))
	draw.Draw(img, img.Bounds(), image.White, image.Point{}, draw.Src)

	// Place the rectangle in the middle of the image, 2px border.
	drawRectangle(img, color.Black, image.Rect(5, 5, 155, 55), 2)

	otf, err := embeddedFont.Scancardium()
	if err != nil {
		return nil, fmt.Errorf("load font: %w", err)
	}

	face, err := opentype.NewFace(otf, &opentype.FaceOptions{
		Size: 20,
		DPI:  72,
	})
	if err != nil {
		return nil, fmt.Errorf("create font face: %w", err)
	}
	defer face.Close()

	fd := font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: face,
	}

	// Center the code inside the image.
	formattedCode := setupCode.Format()
	fd.Dot = fixed.Point26_6{
		X: (fixed.I(img.Bounds().Dx()) - fd.MeasureString(formattedCode)) / 2,
		Y: (fixed.I(img.Bounds().Dy()) + face.Metrics().Ascent) / 2,
	}

	fd.DrawString(formattedCode)

	return img, nil
}

func drawRectangle(img draw.Image, color color.Color, rect image.Rectangle, stroke uint) {
	for i := 0; i < int(stroke); i++ {
		for i := rect.Min.X; i < rect.Max.X; i++ {
			img.Set(i, rect.Min.Y, color)
			img.Set(i, rect.Max.Y, color)
		}
		for i := rect.Min.Y; i <= rect.Max.Y; i++ {
			img.Set(rect.Min.X, i, color)
			img.Set(rect.Max.X, i, color)
		}
		rect = rect.Inset(1)
	}
}
