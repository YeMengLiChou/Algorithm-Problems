package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(countSequences([]int{2, 5}, 2))
	// fmt.Println(countSequences([]int{2, 3, 2}, 6))
	// fmt.Println(countSequences([]int{1, 5}, 1))
	// fmt.Println(countSequences([]int{4, 6, 3}, 2))
	fmt.Println(countSequences([]int{5, 3, 5}, 3))
}

type pair struct {
	a int
	float64
}

func countSequences(nums []int, k int64) int {
	return dfs(0, len(nums), float64(1), float64(k), nums, map[pair]int{})
}

func dfs(idx, n int, sum float64, k float64, nums []int, mp map[pair]int) int {
	if idx == n {

		if math.Abs(sum-k) < 1e-9 {
			return 1
		}
		return 0
	}
	res, ok := mp[pair{idx, sum}]
	if ok {
		return res
	}
	res += dfs(idx+1, n, sum*float64(nums[idx]), k, nums, mp)
	res += dfs(idx+1, n, sum/float64(nums[idx]), k, nums, mp)
	res += dfs(idx+1, n, sum, k, nums, mp)

	mp[pair{idx, sum}] = res

	return res
}
