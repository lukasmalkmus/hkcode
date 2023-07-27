package hk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lukasmalkmus/hkcode/hk"
)

func TestCode(t *testing.T) {
	tests := []struct {
		name       string
		code       hk.Code
		wantString string
		wantFormat string
		wantValid  bool
	}{
		{
			name:       "valid",
			code:       12345678,
			wantString: "12345678",
			wantFormat: "123-45-678",
			wantValid:  true,
		},
		{
			name:       "valid - min",
			code:       0,
			wantString: "00000000",
			wantFormat: "000-00-000",
			wantValid:  true,
		},
		{
			name:       "valid - max",
			code:       99999999,
			wantString: "99999999",
			wantFormat: "999-99-999",
			wantValid:  true,
		},
		{
			name:       "invalid - too long",
			code:       100000000,
			wantString: "100000000",
			wantFormat: "",
			wantValid:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantString, tt.code.String())
			assert.Equal(t, tt.wantFormat, tt.code.Format())
			assert.Equal(t, tt.wantValid, tt.code.Valid())
		})
	}
}
