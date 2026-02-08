package main

import (
	"slices"
)

func main() {
	in := []int{2, 1, 5}
	res := minRemoval(in, 2)
}

func minRemoval(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	res := n - 1
	slices.Sort(nums)
	i := 0
	kk := int64(k)
	for i < n {
		mx := kk * int64(nums[i])
		idx, _ := slices.BinarySearchFunc(nums, mx+1, func(a int, b int64) int {
			return int(int64(a) - b)
		})
		res = min(res, i+n-idx)
		l := i
		for i > 0 && i < n && nums[i-1] == nums[i] {
			i++
		}
		if l == i {
			i++
		}
	}
	return res
}
