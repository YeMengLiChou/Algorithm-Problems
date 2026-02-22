package main

func scoreDifference(nums []int) int {
	A, B := 0, 0
	a, b := &A, &B
	for i, x := range nums {
		if x&1 == 1 {
			a, b = b, a
		}
		if i%6 == 5 {
			a, b = b, a
		}
		*a += nums[i]
	}
	return A - B
}
