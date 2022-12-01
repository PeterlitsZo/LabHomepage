package util

import "strconv"

func String2Uint(s string) (uint, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	return uint(v), err
}

func Uint2String(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}
