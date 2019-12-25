package util

const asciiMaxValue = 128

func IsAscii(i int64) bool {
	return 0 <= i && i <= asciiMaxValue
}
