package euler004

import (
	"math"
	"strconv"
)

func IsPalindromic(n int32) bool {
	strN := strconv.FormatInt(int64(n), 10)
	byteStrN := []rune(strN)

	for i, j := 0, len(byteStrN)-1; i < j; i, j = i+1, j-1 {
		if byteStrN[i] != byteStrN[j] {
			return false
		}
	}

	return true
}

func LargestPalindromeProduct(n int32) int32 {
	for i := n - 1; i > 101100; i-- {
		if !IsPalindromic(i) {
			continue
		}
		sqrtI := int32(math.Sqrt(float64(i)))
		if sqrtI > 999 {
			continue
		}
		for j := sqrtI; j > 99; j-- {
			if i%j == 0 {
				d := i / j
				if d > 999 {
					break
				}

				if d > 99 {
					return i
				}

			}
		}
	}
	return 101101
}
