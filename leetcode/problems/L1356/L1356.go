package main

import (
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		x, y := arr[i], arr[j]
		xcnt, ycnt := bits.OnesCount32(uint32(x)), bits.OnesCount32(uint32(y))
		if xcnt != ycnt {
			return xcnt < ycnt
		} else {
			return x < y
		}
	})
	return arr
}
