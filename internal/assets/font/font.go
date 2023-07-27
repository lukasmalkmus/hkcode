//go:build darwin

package font

import (
	_ "embed"
	"fmt"

	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var (
	//go:embed Scancardium_2.0.ttf
	scancardium []byte
	//go:embed SF-Mono-Bold.otf
	sfMonoBold []byte
)

// Scancardium returns the Scancardium 2.0 font.
func Scancardium() (*opentype.Font, error) {
	ttf, err := opentype.Parse(scancardium)
	if err != nil {
		return nil, fmt.Errorf("parse font: %w", err)
	} else if fontName, _ := ttf.Name(nil, sfnt.NameIDFull); fontName != "Scancardium" {
		return nil, fmt.Errorf("unexpected font name %q", fontName)
	}
	return ttf, nil
}

// SFMonoBold returns the San Franscisco Mono Bold font (SF Mono Bold).
func SFMonoBold() (*opentype.Font, error) {
	ttf, err := opentype.Parse(sfMonoBold)
	if err != nil {
		return nil, fmt.Errorf("parse font: %w", err)
	} else if fontName, _ := ttf.Name(nil, sfnt.NameIDFull); fontName != "SF Mono Bold" {
		return nil, fmt.Errorf("unexpected font name %q", fontName)
	}
	return ttf, nil
}
