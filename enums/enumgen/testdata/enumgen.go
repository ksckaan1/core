// Code generated by "enumgen.test -test.paniconexit0 -test.timeout=10m0s"; DO NOT EDIT.

package testdata

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync/atomic"

	"cogentcore.org/core/enums"
)

var _FruitsValues = []Fruits{0, 1, 2, 3, 4, 5, 6}

// FruitsN is the highest valid value for type Fruits, plus one.
const FruitsN Fruits = 7

var _FruitsValueMap = map[string]Fruits{`Apple`: 0, `apple`: 0, `Orange`: 1, `orange`: 1, `Peach`: 2, `peach`: 2, `Strawberry`: 3, `strawberry`: 3, `Blackberry`: 4, `blackberry`: 4, `Blueberry`: 5, `blueberry`: 5, `Apricot`: 6, `apricot`: 6}

var _FruitsDescMap = map[Fruits]string{0: ``, 1: ``, 2: ``, 3: ``, 4: ``, 5: ``, 6: ``}

var _FruitsMap = map[Fruits]string{0: `Apple`, 1: `Orange`, 2: `Peach`, 3: `Strawberry`, 4: `Blackberry`, 5: `Blueberry`, 6: `Apricot`}

// String returns the string representation of this Fruits value.
func (i Fruits) String() string { return enums.String(i, _FruitsMap) }

// SetString sets the Fruits value from its string representation,
// and returns an error if the string is invalid.
func (i *Fruits) SetString(s string) error {
	return enums.SetStringLower(i, s, _FruitsValueMap, "Fruits")
}

// Int64 returns the Fruits value as an int64.
func (i Fruits) Int64() int64 { return int64(i) }

// SetInt64 sets the Fruits value from an int64.
func (i *Fruits) SetInt64(in int64) { *i = Fruits(in) }

// Desc returns the description of the Fruits value.
func (i Fruits) Desc() string {
	if str, ok := _FruitsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// FruitsValues returns all possible values for the type Fruits.
func FruitsValues() []Fruits { return _FruitsValues }

// Values returns all possible values for the type Fruits.
func (i Fruits) Values() []enums.Enum {
	res := make([]enums.Enum, len(_FruitsValues))
	for i, d := range _FruitsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a valid option for type Fruits.
func (i Fruits) IsValid() bool {
	_, ok := _FruitsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Fruits) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Fruits) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Fruits.UnmarshalText:", err)
	}
	return nil
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i Fruits) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *Fruits) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if err := i.SetString(s); err != nil {
		log.Println("Fruits.UnmarshalJSON:", err)
	}
	return nil
}

var _FoodsValues = []Foods{7, 8, 9, 10}

// FoodsN is the highest valid value for type Foods, plus one.
const FoodsN Foods = 11

var _FoodsValueMap = map[string]Foods{`Bread`: 7, `Lettuce`: 8, `Cheese`: 9, `Meat`: 10}

var _FoodsDescMap = map[Foods]string{7: ``, 8: ``, 9: ``, 10: ``}

var _FoodsMap = map[Foods]string{7: `Bread`, 8: `Lettuce`, 9: `Cheese`, 10: `Meat`}

// String returns the string representation of this Foods value.
func (i Foods) String() string { return enums.StringExtended[Foods, Fruits](i, _FoodsMap) }

// SetString sets the Foods value from its string representation,
// and returns an error if the string is invalid.
func (i *Foods) SetString(s string) error {
	return enums.SetStringExtended(i, (*Fruits)(i), s, _FoodsValueMap)
}

// Int64 returns the Foods value as an int64.
func (i Foods) Int64() int64 { return int64(i) }

// SetInt64 sets the Foods value from an int64.
func (i *Foods) SetInt64(in int64) { *i = Foods(in) }

// Desc returns the description of the Foods value.
func (i Foods) Desc() string {
	if str, ok := _FoodsDescMap[i]; ok {
		return str
	}
	return Fruits(i).Desc()
}

// FoodsValues returns all possible values for the type Foods.
func FoodsValues() []Foods {
	es := FruitsValues()
	res := make([]Foods, len(es))
	for i, e := range es {
		res[i] = Foods(e)
	}
	res = append(res, _FoodsValues...)
	return res
}

// Values returns all possible values for the type Foods.
func (i Foods) Values() []enums.Enum {
	es := FruitsValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_FoodsValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _FoodsValues {
		res[i+les] = d
	}
	return res
}

// IsValid returns whether the value is a valid option for type Foods.
func (i Foods) IsValid() bool {
	_, ok := _FoodsMap[i]
	if !ok {
		return Fruits(i).IsValid()
	}
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Foods) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Foods) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Foods.UnmarshalText:", err)
	}
	return nil
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i Foods) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *Foods) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if err := i.SetString(s); err != nil {
		log.Println("Foods.UnmarshalJSON:", err)
	}
	return nil
}

var _DaysValues = []Days{-11, -9, -7, -5, -3, -1, 1}

// DaysN is the highest valid value for type Days, plus one.
const DaysN Days = 2

var _DaysValueMap = map[string]Days{`DAY_SATURDAY`: -11, `DAY_FRIDAY`: -9, `DAY_THURSDAY`: -7, `DAY_WEDNESDAY`: -5, `DAY_TUESDAY`: -3, `DAY_MONDAY`: -1, `DAY_SUNDAY`: 1}

var _DaysDescMap = map[Days]string{-11: `Saturday is the seventh day of the week`, -9: `Friday is the sixth day of the week`, -7: `Thursday is the fifth day of the week`, -5: `Wednesday is the fourth day of the week`, -3: `Tuesday is the third day of the week`, -1: `Monday is the second day of the week`, 1: `Sunday is the first day of the week`}

var _DaysMap = map[Days]string{-11: `DAY_SATURDAY`, -9: `DAY_FRIDAY`, -7: `DAY_THURSDAY`, -5: `DAY_WEDNESDAY`, -3: `DAY_TUESDAY`, -1: `DAY_MONDAY`, 1: `DAY_SUNDAY`}

// String returns the string representation of this Days value.
func (i Days) String() string { return enums.String(i, _DaysMap) }

// SetString sets the Days value from its string representation,
// and returns an error if the string is invalid.
func (i *Days) SetString(s string) error { return enums.SetString(i, s, _DaysValueMap, "Days") }

// Int64 returns the Days value as an int64.
func (i Days) Int64() int64 { return int64(i) }

// SetInt64 sets the Days value from an int64.
func (i *Days) SetInt64(in int64) { *i = Days(in) }

// Desc returns the description of the Days value.
func (i Days) Desc() string {
	if str, ok := _DaysDescMap[i]; ok {
		return str
	}
	return i.String()
}

// DaysValues returns all possible values for the type Days.
func DaysValues() []Days { return _DaysValues }

// Values returns all possible values for the type Days.
func (i Days) Values() []enums.Enum {
	res := make([]enums.Enum, len(_DaysValues))
	for i, d := range _DaysValues {
		res[i] = d
	}
	return res
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Days) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Days) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Days.UnmarshalText:", err)
	}
	return nil
}

// MarshalGQL implements the [graphql.Marshaler] interface.
func (i Days) MarshalGQL(w io.Writer) {
	w.Write([]byte(strconv.Quote(i.String())))
}

// UnmarshalGQL implements the [graphql.Unmarshaler] interface.
func (i *Days) UnmarshalGQL(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("Days should be a string, but got a value of type %T instead", value)
	}
	return i.SetString(str)
}

var _StatesValues = []States{1, 3, 5, 7, 9, 11, 13}

// StatesN is the highest valid value for type States, plus one.
const StatesN States = 14

var _StatesValueMap = map[string]States{`enabled`: 1, `not-enabled`: 3, `focused`: 5, `vered`: 7, `currently-being-pressed-by-user`: 9, `actively-focused`: 11, `selected`: 13}

var _StatesDescMap = map[States]string{1: `Enabled indicates the widget is enabled`, 3: `Disabled indicates the widget is disabled`, 5: `Focused indicates the widget has keyboard focus`, 7: `Hovered indicates the widget is being hovered over`, 9: `Active indicates the widget is being interacted with`, 11: `ActivelyFocused indicates the widget has active keyboard focus`, 13: `Selected indicates the widget is selected`}

var _StatesMap = map[States]string{1: `enabled`, 3: `not-enabled`, 5: `focused`, 7: `vered`, 9: `currently-being-pressed-by-user`, 11: `actively-focused`, 13: `selected`}

// String returns the string representation of this States value.
func (i States) String() string { return enums.BitFlagString(i, _StatesValues) }

// BitIndexString returns the string representation of this States value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i States) BitIndexString() string { return enums.String(i, _StatesMap) }

// SetString sets the States value from its string representation,
// and returns an error if the string is invalid.
func (i *States) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the States value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *States) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _StatesValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			return fmt.Errorf("%q is not a valid value for type States", flg)
		}
	}
	return nil
}

// Int64 returns the States value as an int64.
func (i States) Int64() int64 { return int64(i) }

// SetInt64 sets the States value from an int64.
func (i *States) SetInt64(in int64) { *i = States(in) }

// Desc returns the description of the States value.
func (i States) Desc() string {
	if str, ok := _StatesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// StatesValues returns all possible values for the type States.
func StatesValues() []States { return _StatesValues }

// Values returns all possible values for the type States.
func (i States) Values() []enums.Enum {
	res := make([]enums.Enum, len(_StatesValues))
	for i, d := range _StatesValues {
		res[i] = d
	}
	return res
}

// HasFlag returns whether these bit flags have the given bit flag set.
func (i States) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *States) SetFlag(on bool, f ...enums.BitFlag) {
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

// MarshalJSON implements the [json.Marshaler] interface.
func (i States) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *States) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if err := i.SetString(s); err != nil {
		log.Println("States.UnmarshalJSON:", err)
	}
	return nil
}

// Scan implements the [driver.Valuer] interface.
func (i States) Value() (driver.Value, error) {
	return i.String(), nil
}

// Value implements the [sql.Scanner] interface.
func (i *States) Scan(value any) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value for type States: %[1]T(%[1]v)", value)
	}

	return i.SetString(str)
}

var _LanguagesValues = []Languages{6, 10, 14, 18, 22, 26, 30, 34, 38, 42, 46, 50, 54}

// LanguagesN is the highest valid value for type Languages, plus one.
const LanguagesN Languages = 55

var _LanguagesValueMap = map[string]Languages{`Go`: 6, `Python`: 10, `JavaScript`: 14, `Dart`: 18, `Rust`: 22, `Ruby`: 26, `C`: 30, `CPP`: 34, `ObjectiveC`: 38, `Java`: 42, `TypeScript`: 46, `Kotlin`: 50, `Swift`: 54}

var _LanguagesDescMap = map[Languages]string{6: `Go is the best programming language`, 10: ``, 14: `JavaScript is the worst programming language`, 18: ``, 22: ``, 26: ``, 30: ``, 34: ``, 38: ``, 42: ``, 46: ``, 50: ``, 54: ``}

var _LanguagesMap = map[Languages]string{6: `Go`, 10: `Python`, 14: `JavaScript`, 18: `Dart`, 22: `Rust`, 26: `Ruby`, 30: `C`, 34: `CPP`, 38: `ObjectiveC`, 42: `Java`, 46: `TypeScript`, 50: `Kotlin`, 54: `Swift`}

// String returns the string representation of this Languages value.
func (i Languages) String() string { return enums.BitFlagString(i, _LanguagesValues) }

// BitIndexString returns the string representation of this Languages value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i Languages) BitIndexString() string { return enums.String(i, _LanguagesMap) }

// SetString sets the Languages value from its string representation,
// and returns an error if the string is invalid.
func (i *Languages) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the Languages value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *Languages) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _LanguagesValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			return fmt.Errorf("%q is not a valid value for type Languages", flg)
		}
	}
	return nil
}

// Int64 returns the Languages value as an int64.
func (i Languages) Int64() int64 { return int64(i) }

// SetInt64 sets the Languages value from an int64.
func (i *Languages) SetInt64(in int64) { *i = Languages(in) }

// Desc returns the description of the Languages value.
func (i Languages) Desc() string {
	if str, ok := _LanguagesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// LanguagesValues returns all possible values for the type Languages.
func LanguagesValues() []Languages { return _LanguagesValues }

// Values returns all possible values for the type Languages.
func (i Languages) Values() []enums.Enum {
	res := make([]enums.Enum, len(_LanguagesValues))
	for i, d := range _LanguagesValues {
		res[i] = d
	}
	return res
}

// HasFlag returns whether these bit flags have the given bit flag set.
func (i Languages) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *Languages) SetFlag(on bool, f ...enums.BitFlag) {
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
func (i Languages) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Languages) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Languages.UnmarshalText:", err)
	}
	return nil
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i Languages) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *Languages) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if err := i.SetString(s); err != nil {
		log.Println("Languages.UnmarshalJSON:", err)
	}
	return nil
}

var _MoreLanguagesValues = []MoreLanguages{55}

// MoreLanguagesN is the highest valid value for type MoreLanguages, plus one.
const MoreLanguagesN MoreLanguages = 56

var _MoreLanguagesValueMap = map[string]MoreLanguages{`Perl`: 55}

var _MoreLanguagesDescMap = map[MoreLanguages]string{55: ``}

var _MoreLanguagesMap = map[MoreLanguages]string{55: `Perl`}

// String returns the string representation of this MoreLanguages value.
func (i MoreLanguages) String() string {
	return enums.BitFlagStringExtended(i, _MoreLanguagesValues, LanguagesValues())
}

// BitIndexString returns the string representation of this MoreLanguages value
// if it is a bit index value (typically an enum constant), and
// not an actual bit flag value.
func (i MoreLanguages) BitIndexString() string {
	return enums.BitIndexStringExtended[MoreLanguages, Languages](i, _MoreLanguagesMap)
}

// SetString sets the MoreLanguages value from its string representation,
// and returns an error if the string is invalid.
func (i *MoreLanguages) SetString(s string) error { *i = 0; return i.SetStringOr(s) }

// SetStringOr sets the MoreLanguages value from its string representation
// while preserving any bit flags already set, and returns an
// error if the string is invalid.
func (i *MoreLanguages) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _MoreLanguagesValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			err := (*Languages)(i).SetStringOr(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the MoreLanguages value as an int64.
func (i MoreLanguages) Int64() int64 { return int64(i) }

// SetInt64 sets the MoreLanguages value from an int64.
func (i *MoreLanguages) SetInt64(in int64) { *i = MoreLanguages(in) }

// Desc returns the description of the MoreLanguages value.
func (i MoreLanguages) Desc() string {
	if str, ok := _MoreLanguagesDescMap[i]; ok {
		return str
	}
	return Languages(i).Desc()
}

// MoreLanguagesValues returns all possible values for the type MoreLanguages.
func MoreLanguagesValues() []MoreLanguages {
	es := LanguagesValues()
	res := make([]MoreLanguages, len(es))
	for i, e := range es {
		res[i] = MoreLanguages(e)
	}
	res = append(res, _MoreLanguagesValues...)
	return res
}

// Values returns all possible values for the type MoreLanguages.
func (i MoreLanguages) Values() []enums.Enum {
	es := LanguagesValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_MoreLanguagesValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _MoreLanguagesValues {
		res[i+les] = d
	}
	return res
}

// HasFlag returns whether these bit flags have the given bit flag set.
func (i MoreLanguages) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given flags in these flags to the given value.
func (i *MoreLanguages) SetFlag(on bool, f ...enums.BitFlag) {
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
func (i MoreLanguages) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *MoreLanguages) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("MoreLanguages.UnmarshalText:", err)
	}
	return nil
}

// MarshalJSON implements the [json.Marshaler] interface.
func (i MoreLanguages) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (i *MoreLanguages) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if err := i.SetString(s); err != nil {
		log.Println("MoreLanguages.UnmarshalJSON:", err)
	}
	return nil
}
