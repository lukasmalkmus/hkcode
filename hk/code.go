package hk

import (
	"fmt"
	"strconv"
)

// ErrInvalidCode can be returned when the [Code] is not valid.
var ErrInvalidCode = fmt.Errorf("invalid setup code")

// Code represents an Apple HomeKit® setup code.
type Code uint32

// String returns a string representation of the code.
//
// Implements [fmt.Stringer].
func (c Code) String() string {
	s := strconv.Itoa(int(c))

	// Pad with zeros, if necessary.
	if l := len(s); l < 8 {
		for i := 0; i < 8-l; i++ {
			s = "0" + s
		}
	}

	return s
}

// Format returns the code in the Apple preferred format XXX-XX-XXX. If the code
// is not valid, it returns an empty string.
func (c Code) Format() string {
	if !c.Valid() {
		return ""
	}
	s := c.String()
	return fmt.Sprintf("%s-%s-%s", s[0:3], s[3:5], s[5:8])
}

// Valid returns true if the code is valid, false otherwise. Valid codes are
// between and including 0 and 99999999.
func (c Code) Valid() bool {
	return c <= 99999999
}
