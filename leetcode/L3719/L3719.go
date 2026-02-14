package main

func longestBalanced(nums []int) int {
	n := len(nums)
	res := 0
	even, odd := make(map[int]int), make(map[int]int)
	update := func(x int) {
		if x&1 == 0 {
			even[x]++
		} else {
			odd[x]++
		}
	}
	for i := range n {
		clear(even)
		clear(odd)
		update(nums[i])
		for j := i + 1; j < n; j++ {
			update(nums[j])
			if len(even) == len(odd) {
				res = max(res, j-i+1)
			}
		}
	}
	return res
}
