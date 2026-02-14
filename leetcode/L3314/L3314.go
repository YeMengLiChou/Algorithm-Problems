package main

// minBitwiseArray 3314. 构造最小位运算数组 I
func minBitwiseArray_1(nums []int) []int {
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

func minBitwiseArray(nums []int) []int {
	for idx, elem := range nums {
		if elem == 2 {
			nums[idx] = -1
		} else {
			t := ^elem
			nums[idx] = ((t & -t) >> 1) ^ elem
		}
	}
	return nums
}
