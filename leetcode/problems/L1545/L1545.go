package main

import (
	"math"
	"strings"
)

func findKthBitV1(n int, k int) byte {
	s := "0"
	var sb strings.Builder
	sb.Grow(int(math.Pow(float64(2), float64(n)) - 1))
	for range n - 1 {
		sb.Reset()
		sb.WriteString(s)
		sb.WriteString("1")
		for j := len(s) - 1; j >= 0; j-- {
			if s[j] == '0' {
				sb.WriteRune('1')
			} else {
				sb.WriteRune('0')
			}
		}
		s = sb.String()
	}
	return s[k-1]
}

func findKthBitV2(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	length := (1 << n) - 1
	mid := length/2 + 1
	if k == mid {
		return '1'
	} else if k > mid {
		k = mid + mid - k
		return '1' - findKthBitV2(n-1, k) + '0'
	} else {
		return findKthBitV2(n-1, k)
	}
}

func findKthBit(n int, k int) byte {
	return findKthBitV2(n, k)
}
