// Code generated by "stringer -type Code -linecomment ./pkg/errcode"; DO NOT EDIT.

package errcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ParamErr-1000]
	_ = x[SourceNotFind-1001]
	_ = x[SystemErr-1002]
	_ = x[SQLErr-1003]
	_ = x[MarshalErr-1004]
}

const _Code_name = "参数错误资源不存在系统错误数据库错误json错误"

var _Code_index = [...]uint8{0, 12, 27, 39, 54, 64}

func (i Code) String() string {
	i -= 1000
	if i >= Code(len(_Code_index)-1) {
		return "Code(" + strconv.FormatInt(int64(i+1000), 10) + ")"
	}
	return _Code_name[_Code_index[i]:_Code_index[i+1]]
}
