package main

import "sort"

func maximizeSquareHoleArea(_ int, _ int, hBars []int, vBars []int) int {
	ans := min(find(hBars), find(vBars)) + 1
	return ans * ans
}

func find(bars []int) int {
	sort.Ints(bars)
	ans, cnt := 0, 0
	for i, x := range bars {
		if i > 0 && bars[i-1]+1 == x {
			cnt++
		} else {
			cnt = 1
		}
		ans = max(ans, cnt)
	}
	return ans
}
