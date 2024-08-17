package utils

import "strconv"

const (
	Base10    = 10
	BitSize64 = 64
)

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, Base10, BitSize64)
}
