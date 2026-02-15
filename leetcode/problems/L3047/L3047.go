package main

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	ans := 0
	n := len(bottomLeft)
	for i := range n {
		for j := i + 1; j < n; j++ {
			w := min(topRight[i][1], topRight[j][1]) - max(bottomLeft[i][1], bottomLeft[j][1])
			h := min(topRight[i][0], topRight[j][0]) - max(bottomLeft[i][0], bottomLeft[j][0])
			ans = max(ans, min(w, h))
		}
	}
	return int64(ans) * int64(ans)
}
