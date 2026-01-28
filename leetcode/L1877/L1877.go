package main

import (
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	ans, size := 0, len(nums)
	i, j := 0, size-1

	for i < j {
		ans = max(ans, nums[i]+nums[j])
		i++
		j--
	}
	return ans
}
