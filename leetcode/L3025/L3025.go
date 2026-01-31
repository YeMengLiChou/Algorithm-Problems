package main

import "fmt"

func main() {
	points := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	ans := numberOfPairs(points)
	fmt.Println(ans)
}

func numberOfPairs(points [][]int) int {
	n := len(points)
	ans := 0
	for i := range n {
		for j := range n {
			if i == j {
				continue
			}
			// not left top
			if points[i][0] > points[j][0] || points[i][1] < points[j][1] {
				continue
			}

			lx := min(points[i][0], points[j][0])
			rx := max(points[i][0], points[j][0])
			ly := min(points[i][1], points[j][1])
			ry := max(points[i][1], points[j][1])

			// check if contains any points
			valid := true
			for k, p := range points {
				if i == k || j == k {
					continue
				}
				x, y := p[0], p[1]
				if lx <= x && x <= rx && ly <= y && y <= ry {
					// contains this point
					valid = false
					break
				}
			}
			if valid {
				ans++
			}
		}
	}
	return ans
}
