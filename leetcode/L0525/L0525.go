package main

func findMaxLengthV1(nums []int) int {
	n := len(nums)
	s0, s1 := make([]int, n+1), make([]int, n+1)
	for i, x := range nums {
		if x&1 == 0 {
			s0[i+1] = 1
		} else {
			s1[i+1] = 1
		}
		s0[i+1] += s0[i]
		s1[i+1] += s1[i]
	}
	res := 0
	mp := make(map[int]int)
	for i := range n + 1 {
		s := s1[i] - s0[i]
		pre, ok := mp[s]
		if ok {
			res = max(res, i-pre)
		} else {
			mp[s] = i
		}
	}
	return res
}

func findMaxLength(nums []int) int {
	return findMaxLengthV1(nums)
}
