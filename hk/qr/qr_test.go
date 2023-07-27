package qr_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lukasmalkmus/hkcode/hk"
	"github.com/lukasmalkmus/hkcode/hk/qr"
	"github.com/lukasmalkmus/hkcode/internal/testdata"
	"github.com/lukasmalkmus/hkcode/internal/testutil"
)

func TestCreateCode(t *testing.T) {
	golden := testdata.GetGoldenQRCodeImage(t)

	img, err := qr.CreateCode(12344321, "RFGD", hk.FlagIP|hk.FlagBTLE, hk.CategorySwitch)
	require.NoError(t, err)

	testutil.AssertEqualImage(t, golden, img)
}

func TestCreateBoxedCode(t *testing.T) {
	golden := testdata.GetGoldenBoxedQRCodeImage(t)

	img, err := qr.CreateBoxedCode(12345678, "RFGD", hk.FlagIP|hk.FlagBTLE, hk.CategorySwitch)
	require.NoError(t, err)

	testutil.AssertEqualImage(t, golden, img)
}
