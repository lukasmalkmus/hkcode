package hk

import (
	"fmt"
	"strings"
)

// Flag represents an Apple HomeKit® setup flag. Setup flags indicate the
// supported pairing methods. However, Apple devices seem to ignore them and it
// doesn't matter which flags are set, if even.
type Flag uint8

// All available setup flags.
const (
	FlagNone Flag = 0         // <none>
	FlagNFC  Flag = 1 << iota // NFC
	FlagIP                    // IP
	FlagBTLE                  // BTLE
	maxFlag
)

// String returns a string representation of the flag.
//
// It implements [fmt.Stringer].
func (f Flag) String() string {
	if f >= maxFlag {
		return fmt.Sprintf("<unknown field type: %d (%08b)>", f, f)
	}

	switch f {
	case FlagNFC:
		return "NFC"
	case FlagIP:
		return "IP"
	case FlagBTLE:
		return "BTLE"
	}

	var res []string
	for flag := FlagNFC; flag < maxFlag; flag <<= 1 {
		if f&flag != 0 {
			res = append(res, flag.String())
		}
	}
	return strings.Join(res, "|")
}
