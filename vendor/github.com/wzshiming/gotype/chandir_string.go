// Code generated by "stringer -type ChanDir chandir.go"; DO NOT EDIT.

package gotype

import "strconv"

const _ChanDir_name = "RecvDirSendDir"

var _ChanDir_index = [...]uint8{0, 7, 14}

func (i ChanDir) String() string {
	i -= 1
	if i < 0 || i >= ChanDir(len(_ChanDir_index)-1) {
		return "ChanDir(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ChanDir_name[_ChanDir_index[i]:_ChanDir_index[i+1]]
}
