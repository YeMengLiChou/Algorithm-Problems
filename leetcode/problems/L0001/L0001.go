package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, x := range nums {
		l, ok := mp[x]
		if !ok {
			mp[target-x] = i
		} else {
			return []int{l, i}
		}
	}
	return []int{}
}
