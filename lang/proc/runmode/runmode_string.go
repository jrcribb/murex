// Code generated by "stringer -type=RunMode"; DO NOT EDIT.

package runmode

import "fmt"

const _RunMode_name = "NormalTryTryPipeEvil"

var _RunMode_index = [...]uint8{0, 6, 9, 16, 20}

func (i RunMode) String() string {
	if i < 0 || i >= RunMode(len(_RunMode_index)-1) {
		return fmt.Sprintf("RunMode(%d)", i)
	}
	return _RunMode_name[_RunMode_index[i]:_RunMode_index[i+1]]
}