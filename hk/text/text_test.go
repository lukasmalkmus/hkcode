package text_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lukasmalkmus/hkcode/hk/text"
	"github.com/lukasmalkmus/hkcode/internal/testdata"
	"github.com/lukasmalkmus/hkcode/internal/testutil"
)

func TestCreateCode(t *testing.T) {
	golden := testdata.GetGoldenTextImage(t)

	img, err := text.CreateCode(12344321)
	require.NoError(t, err)

	testutil.AssertEqualImage(t, golden, img)
}
