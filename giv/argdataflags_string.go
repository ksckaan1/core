// Code generated by "stringer -type=ArgDataFlags"; DO NOT EDIT.

package giv

import (
	"fmt"
	"strconv"
)

const _ArgDataFlags_name = "ArgDataHasDefArgDataValSetArgDataFlagsN"

var _ArgDataFlags_index = [...]uint8{0, 13, 26, 39}

func (i ArgDataFlags) String() string {
	if i < 0 || i >= ArgDataFlags(len(_ArgDataFlags_index)-1) {
		return "ArgDataFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ArgDataFlags_name[_ArgDataFlags_index[i]:_ArgDataFlags_index[i+1]]
}

func (i *ArgDataFlags) FromString(s string) error {
	for j := 0; j < len(_ArgDataFlags_index)-1; j++ {
		if s == _ArgDataFlags_name[_ArgDataFlags_index[j]:_ArgDataFlags_index[j+1]] {
			*i = ArgDataFlags(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type ArgDataFlags", s)
}
