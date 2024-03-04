// Code generated by "core generate"; DO NOT EDIT.

package xyzv

import (
	"errors"
	"log"
	"strconv"

	"cogentcore.org/core/enums"
)

var _SelModesValues = []SelModes{0, 1, 2, 3}

// SelModesN is the highest valid value for type SelModes, plus one.
const SelModesN SelModes = 4

var _SelModesNameToValueMap = map[string]SelModes{`NotSelectable`: 0, `Selectable`: 1, `SelectionBox`: 2, `Manipulable`: 3}

var _SelModesDescMap = map[SelModes]string{0: `NotSelectable means that selection events are ignored entirely`, 1: `Selectable means that nodes can be selected but no visible consequence occurs`, 2: `SelectionBox means that a selection bounding box is drawn around selected nodes`, 3: `Manipulable means that a manipulation box will be created for selected nodes, which can update the Pose parameters dynamically.`}

var _SelModesMap = map[SelModes]string{0: `NotSelectable`, 1: `Selectable`, 2: `SelectionBox`, 3: `Manipulable`}

// String returns the string representation of this SelModes value.
func (i SelModes) String() string {
	if str, ok := _SelModesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the SelModes value from its string representation,
// and returns an error if the string is invalid.
func (i *SelModes) SetString(s string) error {
	if val, ok := _SelModesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type SelModes")
}

// Int64 returns the SelModes value as an int64.
func (i SelModes) Int64() int64 { return int64(i) }

// SetInt64 sets the SelModes value from an int64.
func (i *SelModes) SetInt64(in int64) { *i = SelModes(in) }

// Desc returns the description of the SelModes value.
func (i SelModes) Desc() string {
	if str, ok := _SelModesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// SelModesValues returns all possible values for the type SelModes.
func SelModesValues() []SelModes { return _SelModesValues }

// Values returns all possible values for the type SelModes.
func (i SelModes) Values() []enums.Enum {
	res := make([]enums.Enum, len(_SelModesValues))
	for i, d := range _SelModesValues {
		res[i] = d
	}
	return res
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i SelModes) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *SelModes) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("SelModes.UnmarshalText:", err)
	}
	return nil
}
