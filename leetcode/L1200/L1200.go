package main

import (
	"math"
	"sort"
)

// 解法1: 存一下每个差值的下标
// func minimumAbsDifference(arr []int) [][]int {
// 	sort.Ints(arr)
// 	diff := make(map[int][]int)
// 	size, minD := len(arr), math.MaxInt32

// 	for i := 1; i < size; i++ {
// 		d := arr[i] - arr[i-1]
// 		values, ok := diff[d]
// 		if !ok {
// 			values = make([]int, 0)
// 		}
// 		values = append(values, i)
// 		diff[d] = values
// 		minD = min(minD, d)
// 	}
// 	values, _ := diff[minD]
// 	ans := make([][]int, len(values))
// 	for idx, value := range values {
// 		elem := append(make([]int, 0), arr[value-1], arr[value])
// 		ans[idx] = elem
// 	}
// 	return ans
// }

// 优化:不需要存储其他

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	size, minD := len(arr), math.MaxInt32

	for i := 1; i < size; i++ {
		d := arr[i] - arr[i-1]
		minD = min(minD, d)
	}
	ans := make([][]int, 0)
	for i := 1; i < size; i++ {
		if minD == arr[i]-arr[i-1] {
			ans = append(ans, []int{
				arr[i-1], arr[i],
			})
		}
	}

	return ans
}
