package main

import (
	"math"
	"slices"
)

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	check := func(x int64) bool {
		var total int64
		for _, w := range workerTimes {
			total += int64(math.Sqrt(1+8*float64(x)/float64(w))-1) / 2
		}
		return total >= int64(mountainHeight)
	}

	mx := slices.Max(workerTimes)
	var (
		left  int64 = 0
		right int64 = int64(mountainHeight+1) * int64(mountainHeight) / 2 * int64(mx)
	)
	for left < right {
		mid := (left + right + 1) >> 1
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

