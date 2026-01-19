package main

// minBitwiseArray 3314. 构造最小位运算数组 I
func minBitwiseArray(nums []int) []int {
	find := func(x int) int {
		if x == 2 {
			return -1
		}
		d := 1
		for (x>>d)&1 == 1 {
			d += 1
		}
		return x & (^(1 << (d - 1)))
	}

	for idx, elem := range nums {
		nums[idx] = find(elem)
	}
	return nums
}
