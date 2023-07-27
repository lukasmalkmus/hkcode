package hk

import (
	"fmt"
	"strings"
)

// ErrInvalidID can be returned when the [ID] is not valid.
var ErrInvalidID = fmt.Errorf("invalid setup id")

// ID represents an Apple HomeKit® setup id.
type ID string

// String returns a string representation of the code. All characters are
// uppercased.
//
// Implements [fmt.Stringer].
func (id ID) String() string {
	return strings.ToUpper(string(id))
}

// Valid returns true if the id is valid, false otherwise. Valid ids have 4
// digits.
func (id ID) Valid() bool {
	return len(id) == 4
}
