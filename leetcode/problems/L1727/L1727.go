package main

import (
	"slices"
)

// 思路：枚举每一行的连续1长度，然后排序
func largestSubmatrix(matrix [][]int) int {
	n := len(matrix[0])
	heights := make([]int, n)
	res := 0
	for _, row := range matrix {
		for j, x := range row {
			if x == 0 {
				heights[j] = 0
			} else {
				heights[j] += 1
			}
		}
		tmp := slices.Clone(heights)
		slices.Sort(tmp)
		for j, h := range tmp {
			res = max(res, (n-j)*h)
		}
	}
	return res
}

func largestSubmatrixV1(matrix [][]int) int {
	var (
		n                        = len(matrix[0])
		res, left, right         = 0, 0, 0
		heights, zeros, nonZeros = make([]int, n), make([]int, n), make([]int, n)
		indices                  = make([]int, n)
	)

	for _, row := range matrix {
		left, right = 0, 0
		for _, x := range indices {
			if row[x] == 0 {
				zeros[left] = x
				left++
				heights[x] = 0
			} else {
				nonZeros[right] = x
				right++
				heights[x] += 1
			}
		}
		idx := 0
		for i := range left {
			indices[idx] = zeros[i]
			idx++
		}
		for i := range right {
			indices[idx] = nonZeros[i]
			idx++
		}
		for i, x := range indices {
			res = max(res, heights[x]*(n-i))
		}

	}
	return res
}
