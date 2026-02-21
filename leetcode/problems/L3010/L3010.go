package main

import "math"

// 暴力枚举
func minimumCost1(nums []int) int {
	res := math.MaxInt
	n := len(nums)
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res = min(res, nums[i]+nums[j]+nums[0])
		}
	}
	return res
}

func minimumCost(nums []int) int {
	n := len(nums)
	minV, subMinV := math.MaxInt, math.MaxInt
	for i := 1; i < n; i++ {
		if minV > nums[i] {
			subMinV = minV
			minV = nums[i]
		} else if subMinV > nums[i] {
			subMinV = nums[i]
		}
	}
	return nums[0] + minV + subMinV
}
