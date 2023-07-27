package hk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lukasmalkmus/hkcode/hk"
)

func TestID(t *testing.T) {
	tests := []struct {
		name       string
		id         hk.ID
		wantString string
		wantValid  bool
	}{
		{
			name:       "valid - digits only",
			id:         "1234",
			wantString: "1234",
			wantValid:  true,
		},
		{
			name:       "valid - letters only",
			id:         "abcd",
			wantString: "ABCD",
			wantValid:  true,
		},
		{
			name:       "valid - mixed digits and letters",
			id:         "a1b2",
			wantString: "A1B2",
			wantValid:  true,
		},
		{
			name:       "invalid - too short",
			id:         "123",
			wantString: "123",
			wantValid:  false,
		},
		{
			name:       "invalid - too long",
			id:         "abcde",
			wantString: "ABCDE",
			wantValid:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantString, tt.id.String())
			assert.Equal(t, tt.wantValid, tt.id.Valid())
		})
	}
}
