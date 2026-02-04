package main

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i, x := range nums {
		switch {
		case x == 0:
			res[i] = x
		case x > 0:
			next := i + x%n
			if next < n {
				res[i] = nums[next]
			} else {
				res[i] = nums[next-n]
			}
		case x < 0:
			next := i + x%n
			if next >= 0 {
				res[i] = nums[next]
			} else {
				res[i] = nums[n+next]
			}
		}
	}
	return res
}
