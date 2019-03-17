package my_sol_1

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	s := strconv.FormatInt(int64(x), 10)
	var isNegative bool
	if s[0] == '-' {
		isNegative = true
		s = s[1:]
	}

	// 字串中的每一個字元都可以對應一個byte
	s = reverseString([]byte(s))

	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	if isNegative {
		n = n * -1
	}

	maxInt32 := int32(^uint32(0) >> 1)
	minInt32 := ^maxInt32

	maxValue := int64(maxInt32)
	minValue := int64(minInt32)

	if n > maxValue || n < minValue {
		return 0
	}

	return int(n)
}

// 字串中的一個字元對應一個byte
// string 和 []byte 可以互相轉換
func reverseString(b []byte) string {
	for i, j := 0, len(b)-1; i <= j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
