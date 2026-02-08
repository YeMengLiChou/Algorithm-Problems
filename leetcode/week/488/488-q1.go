package main

func dominantIndices(nums []int) int {
	n := len(nums)
	sum := nums[n-1]
	cnt := 1
	res := 0
	for i := n - 2; i >= 0; i-- {
		if nums[i] > sum/cnt {
			res++
		}
		sum += nums[i]
		cnt++
	}
	return res
}
