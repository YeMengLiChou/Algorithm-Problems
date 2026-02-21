package main

import "math"

func maxSubArrayV1(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}
	mn := math.MaxInt >> 1
	maxSum := math.MinInt

	for _, x := range sum {
		maxSum = max(maxSum, x-mn)
		mn = min(mn, x)
	}
	return maxSum
}
