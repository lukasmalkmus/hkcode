package hk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlag_String(t *testing.T) {
	assert.Equal(t, FlagNone, Flag(0))

	typ := FlagIP
	assert.Equal(t, "IP", typ.String())

	typ |= FlagBTLE
	assert.Equal(t, "IP|BTLE", typ.String())
}
