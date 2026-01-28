package main

import (
	"math"
	"sort"
)


func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)
	i, j, n := 0, k - 1, len(nums)
	ans := math.MaxInt32
	for j < n {
		ans = min(ans, nums[j] - nums[i])
		i ++
		j ++
	}
	return ans
}
