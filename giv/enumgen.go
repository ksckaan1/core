// Code generated by "core generate"; DO NOT EDIT.

package giv

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync/atomic"

	"cogentcore.org/core/enums"
	"cogentcore.org/core/gi"
)

var _SliceViewFlagsValues = []SliceViewFlags{8, 9, 10, 11, 12, 13, 14, 15}

// SliceViewFlagsN is the highest valid value for type SliceViewFlags, plus one.
const SliceViewFlagsN SliceViewFlags = 16

var _SliceViewFlagsNameToValueMap = map[string]SliceViewFlags{`Configured`: 8, `IsArray`: 9, `ShowIndex`: 10, `ReadOnlyKeyNav`: 11, `SelectMode`: 12, `ReadOnlyMultiSel`: 13, `InFocusGrab`: 14, `InFullRebuild`: 15}

var _SliceViewFlagsDescMap = map[SliceViewFlags]string{8: `SliceViewConfigured indicates that the widgets have been configured`, 9: `SliceViewIsArray is whether the slice is actually an array -- no modifications -- set by SetSlice`, 10: `SliceViewShowIndex is whether to show index or not`, 11: `SliceViewReadOnlyKeyNav is whether support key navigation when ReadOnly (default true). uses a capture of up / down events to manipulate selection, not focus.`, 12: `SliceViewSelectMode is whether to be in select rows mode or editing mode`, 13: `SliceViewReadOnlyMultiSel: if view is ReadOnly, default selection mode is to choose one row only. If this is true, standard multiple selection logic with modifier keys is instead supported`, 14: `SliceViewInFocusGrab is a guard for recursive focus grabbing`, 15: `SliceViewInFullRebuild is a guard for recursive rebuild`}

var _SliceViewFlagsMap = map[SliceViewFlags]string{8: `Configured`, 9: `IsArray`, 10: `ShowIndex`, 11: `ReadOnlyKeyNav`, 12: `SelectMode`, 13: `ReadOnlyMultiSel`, 14: `InFocusGrab`, 15: `InFullRebuild`}

// String returns the string representation of this SliceViewFlags value.
func (i SliceViewFlags) String() string {
	str := ""
	for _, ie := range gi.WidgetFlagsValues() {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	for _, ie := range _SliceViewFlagsValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string representation of this SliceViewFlags value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i SliceViewFlags) BitIndexString() string {
	if str, ok := _SliceViewFlagsMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).BitIndexString()
}

// SetString sets the SliceViewFlags value from its string representation,
// and returns an error if the string is invalid.
func (i *SliceViewFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the SliceViewFlags value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *SliceViewFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _SliceViewFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			err := (*gi.WidgetFlags)(i).SetStringOr(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the SliceViewFlags value as an int64.
func (i SliceViewFlags) Int64() int64 { return int64(i) }

// SetInt64 sets the SliceViewFlags value from an int64.
func (i *SliceViewFlags) SetInt64(in int64) { *i = SliceViewFlags(in) }

// Desc returns the description of the SliceViewFlags value.
func (i SliceViewFlags) Desc() string {
	if str, ok := _SliceViewFlagsDescMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).Desc()
}

// SliceViewFlagsValues returns all possible values for the type SliceViewFlags.
func SliceViewFlagsValues() []SliceViewFlags {
	es := gi.WidgetFlagsValues()
	res := make([]SliceViewFlags, len(es))
	for i, e := range es {
		res[i] = SliceViewFlags(e)
	}
	res = append(res, _SliceViewFlagsValues...)
	return res
}

// Values returns all possible values for the type SliceViewFlags.
func (i SliceViewFlags) Values() []enums.Enum {
	es := gi.WidgetFlagsValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_SliceViewFlagsValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _SliceViewFlagsValues {
		res[i+les] = d
	}
	return res
}

// HasFlag returns whether these bit flags have the given bit flag set.
func (i SliceViewFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *SliceViewFlags) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i SliceViewFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *SliceViewFlags) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("SliceViewFlags.UnmarshalText:", err)
	}
	return nil
}

var _TreeViewFlagsValues = []TreeViewFlags{8, 9, 10}

// TreeViewFlagsN is the highest valid value for type TreeViewFlags, plus one.
const TreeViewFlagsN TreeViewFlags = 11

var _TreeViewFlagsNameToValueMap = map[string]TreeViewFlags{`Closed`: 8, `SelectMode`: 9, `TreeViewInOpen`: 10}

var _TreeViewFlagsDescMap = map[TreeViewFlags]string{8: `TreeViewFlagClosed means node is toggled closed (children not visible) Otherwise Open.`, 9: `TreeViewFlagSelectMode, when set on the Root node, determines whether keyboard movements update selection or not.`, 10: `TreeViewInOpen is set in the Open method to prevent recursive opening for lazy-open nodes`}

var _TreeViewFlagsMap = map[TreeViewFlags]string{8: `Closed`, 9: `SelectMode`, 10: `TreeViewInOpen`}

// String returns the string representation of this TreeViewFlags value.
func (i TreeViewFlags) String() string {
	str := ""
	for _, ie := range gi.WidgetFlagsValues() {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	for _, ie := range _TreeViewFlagsValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string representation of this TreeViewFlags value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i TreeViewFlags) BitIndexString() string {
	if str, ok := _TreeViewFlagsMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).BitIndexString()
}

// SetString sets the TreeViewFlags value from its string representation,
// and returns an error if the string is invalid.
func (i *TreeViewFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the TreeViewFlags value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *TreeViewFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _TreeViewFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			err := (*gi.WidgetFlags)(i).SetStringOr(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the TreeViewFlags value as an int64.
func (i TreeViewFlags) Int64() int64 { return int64(i) }

// SetInt64 sets the TreeViewFlags value from an int64.
func (i *TreeViewFlags) SetInt64(in int64) { *i = TreeViewFlags(in) }

// Desc returns the description of the TreeViewFlags value.
func (i TreeViewFlags) Desc() string {
	if str, ok := _TreeViewFlagsDescMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).Desc()
}

// TreeViewFlagsValues returns all possible values for the type TreeViewFlags.
func TreeViewFlagsValues() []TreeViewFlags {
	es := gi.WidgetFlagsValues()
	res := make([]TreeViewFlags, len(es))
	for i, e := range es {
		res[i] = TreeViewFlags(e)
	}
	res = append(res, _TreeViewFlagsValues...)
	return res
}

// Values returns all possible values for the type TreeViewFlags.
func (i TreeViewFlags) Values() []enums.Enum {
	es := gi.WidgetFlagsValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_TreeViewFlagsValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _TreeViewFlagsValues {
		res[i+les] = d
	}
	return res
}

// HasFlag returns whether these bit flags have the given bit flag set.
func (i TreeViewFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *TreeViewFlags) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i TreeViewFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *TreeViewFlags) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("TreeViewFlags.UnmarshalText:", err)
	}
	return nil
}

var _ValueFlagsValues = []ValueFlags{0, 1, 2, 3, 4}

// ValueFlagsN is the highest valid value for type ValueFlags, plus one.
const ValueFlagsN ValueFlags = 5

var _ValueFlagsNameToValueMap = map[string]ValueFlags{`ReadOnly`: 0, `MapKey`: 1, `HasSavedLabel`: 2, `HasSavedDoc`: 3, `DialogNewWindow`: 4}

var _ValueFlagsDescMap = map[ValueFlags]string{0: `ValueReadOnly flagged after first configuration`, 1: `ValueMapKey for OwnKind = Map, this value represents the Key -- otherwise the Value`, 2: `ValueHasSavedLabel is whether the value has a saved version of its label, which can be set either automatically or explicitly`, 3: `ValueHasSavedDoc is whether the value has a saved version of its documentation, which can be set either automatically or explicitly`, 4: `ValueDialogNewWindow indicates that the dialog should be opened with in a new window, instead of a typical FullWindow in same current window. this is triggered by holding down any modifier key while clicking on a button that opens the window.`}

var _ValueFlagsMap = map[ValueFlags]string{0: `ReadOnly`, 1: `MapKey`, 2: `HasSavedLabel`, 3: `HasSavedDoc`, 4: `DialogNewWindow`}

// String returns the string representation of this ValueFlags value.
func (i ValueFlags) String() string {
	str := ""
	for _, ie := range _ValueFlagsValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string representation of this ValueFlags value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i ValueFlags) BitIndexString() string {
	if str, ok := _ValueFlagsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the ValueFlags value from its string representation,
// and returns an error if the string is invalid.
func (i *ValueFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the ValueFlags value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *ValueFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _ValueFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			return fmt.Errorf("%q is not a valid value for type ValueFlags", flg)
		}
	}
	return nil
}

// Int64 returns the ValueFlags value as an int64.
func (i ValueFlags) Int64() int64 { return int64(i) }

// SetInt64 sets the ValueFlags value from an int64.
func (i *ValueFlags) SetInt64(in int64) { *i = ValueFlags(in) }

// Desc returns the description of the ValueFlags value.
func (i ValueFlags) Desc() string {
	if str, ok := _ValueFlagsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ValueFlagsValues returns all possible values for the type ValueFlags.
func ValueFlagsValues() []ValueFlags { return _ValueFlagsValues }

// Values returns all possible values for the type ValueFlags.
func (i ValueFlags) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ValueFlagsValues))
	for i, d := range _ValueFlagsValues {
		res[i] = d
	}
	return res
}

// HasFlag returns whether these bit flags have the given bit flag set.
func (i ValueFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *ValueFlags) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i ValueFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *ValueFlags) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("ValueFlags.UnmarshalText:", err)
	}
	return nil
}
