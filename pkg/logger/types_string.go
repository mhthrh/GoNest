// Code generated by "stringer -type=Types"; DO NOT EDIT.

package logger

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[none-0]
	_ = x[info-1]
	_ = x[warn-2]
	_ = x[typeError-3]
	_ = x[fatal-4]
	_ = x[typePanic-5]
}

const _Types_name = "noneinfowarntypeErrorfataltypePanic"

var _Types_index = [...]uint8{0, 4, 8, 12, 21, 26, 35}

func (i Types) String() string {
	if i < 0 || i >= Types(len(_Types_index)-1) {
		return "Types(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Types_name[_Types_index[i]:_Types_index[i+1]]
}
