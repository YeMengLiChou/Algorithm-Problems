package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n, m := len(nums1), len(nums2)
	stack := make([]int, m)
	cnt := 0
	mp := make(map[int]int)
	res := make([]int, n)
	// 单调栈，单调递增
	for i := m - 1; i >= 0; i-- {
		x := nums2[i]
		for cnt > 0 && nums2[stack[cnt-1]] <= x {
			cnt--
		}
		if cnt > 0 {
			mp[x] = nums2[stack[cnt-1]]
		} else {
			mp[x] = -1
		}
		stack[cnt] = i
		cnt++
	}
	for i := range n {
		res[i] = mp[nums1[i]]
	}
	return res
}
