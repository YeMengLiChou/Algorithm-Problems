package main

import "slices"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i, row := range matrix {
		for j := i + 1; j < n; j++ {
			row[j], matrix[j][i] = matrix[j][i], row[j]
		}
		slices.Reverse(row)
	}
}
