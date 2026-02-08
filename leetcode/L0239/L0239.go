package main

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	dq, hh, tt := make([]int, n), 0, 0
	res := make([]int, n-k+1)
	cnt := 0

	// 单调队列，单调递减
	for i := range k {
		x := nums[i]
		for tt > hh && nums[dq[tt-1]] <= x {
			tt--
		}
		dq[tt] = i
		tt++
	}
	res[cnt] = nums[dq[hh]]
	cnt++
	for i := k; i < n; i++ {
		if i-k == dq[hh] {
			hh++
		}
		x := nums[i]
		for tt > hh && nums[dq[tt-1]] <= x {
			tt--
		}
		dq[tt] = i
		tt++
		res[cnt] = nums[dq[hh]]
		cnt++
	}
	return res
}
