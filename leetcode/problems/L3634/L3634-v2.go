package main

import (
	"slices"
)

func minRemoval(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	var (
		res = n - 1
		kk  = int64(k)
	)
	slices.Sort(nums)

	for l, r := 0, 0; l < n; l++ {
		mx := int64(nums[l]) * kk
		for r < n && int64(nums[r]) <= mx {
			r++
		}
		res = min(res, l+n-r)
	}
	return res
}
