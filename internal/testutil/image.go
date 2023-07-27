package testutil

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertEqualImage(t *testing.T, a, b image.Image) {
	t.Helper()

	ba := a.Bounds()
	bb := b.Bounds()
	if ba != bb {
		t.Errorf("image bounds not equal: %+v, %+v", a.Bounds(), b.Bounds())
		return
	}

	var accumError int64
	diffImg := image.NewRGBA(image.Rect(
		ba.Min.X,
		ba.Min.Y,
		ba.Max.X,
		ba.Max.Y,
	))
	draw.Draw(diffImg, diffImg.Bounds(), a, image.Point{0, 0}, draw.Src)

	for x := ba.Min.X; x < ba.Max.X; x++ {
		for y := ba.Min.Y; y < ba.Max.Y; y++ {
			r1, g1, b1, a1 := a.At(x, y).RGBA()
			r2, g2, b2, a2 := b.At(x, y).RGBA()

			diff := int64(sqDiffUInt32(r1, r2))
			diff += int64(sqDiffUInt32(g1, g2))
			diff += int64(sqDiffUInt32(b1, b2))
			diff += int64(sqDiffUInt32(a1, a2))

			if diff > 0 {
				accumError += diff
				diffImg.Set(
					ba.Min.X+x,
					ba.Min.Y+y,
					color.RGBA{R: 255, A: 255})
			}
		}
	}

	if accumError = int64(math.Sqrt(float64(accumError))); accumError > 0 {
		t.Errorf("images not equal, accumulated error: %d", accumError)

		var buf bytes.Buffer
		if assert.NoError(t, png.Encode(&buf, diffImg)) {
			diffImgPath := path.Join(os.TempDir(), "diff.png")
			if assert.NoError(t, os.WriteFile(diffImgPath, buf.Bytes(), 0666)) {
				t.Logf("image diff written to %s", diffImgPath)
			}
		}
	}
}

func sqDiffUInt32(x, y uint32) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}
