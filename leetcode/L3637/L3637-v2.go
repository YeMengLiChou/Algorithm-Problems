package main

func isTrionic(nums []int) bool {
	if nums[0] >= nums[1] {
		return false
	}
	var (
		n   = len(nums)
		cnt = 0
	)

	for i := 2; i < n; i++ {
		if nums[i-1] == nums[i] {
			return false
		}
		if (nums[i-1]-nums[i-2])*(nums[i]-nums[i-1]) < 0 {
			cnt++
		}
	}
	return cnt == 2
}
