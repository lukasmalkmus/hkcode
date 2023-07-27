package assets

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
)

//go:embed box.png
var box []byte

// Box returns the setup code box template image.
func Box() (image.Image, error) {
	return png.Decode(bytes.NewReader(box))
}
