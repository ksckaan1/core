// Code generated by "stringer -type=Actions"; DO NOT EDIT.

package parse

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

const _Actions_name = "ChgTokenAddSymbolPushScopePushNewScopePopScopeAddDetailActionsN"

var _Actions_index = [...]uint8{0, 8, 17, 26, 38, 46, 55, 63}

func (i Actions) String() string {
	if i < 0 || i >= Actions(len(_Actions_index)-1) {
		return "Actions(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Actions_name[_Actions_index[i]:_Actions_index[i+1]]
}

func (i *Actions) FromString(s string) error {
	for j := 0; j < len(_Actions_index)-1; j++ {
		if s == _Actions_name[_Actions_index[j]:_Actions_index[j+1]] {
			*i = Actions(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Actions")
}
