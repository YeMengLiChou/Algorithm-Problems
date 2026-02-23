package main

import "fmt"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	f := make([]float64, query_row+1)
	f[0] = float64(poured)
	for i := range query_row {
		origin := float64(0)
		for j := 0; j <= i; j++ {
			rest := f[j] - 1
			if rest > 0 {
				f[j] = origin + rest/2
				origin = f[j+1] // 记录原本的下一个
				f[j+1] = rest / 2
			}
		}
		fmt.Println(f)
	}
	return f[query_glass-1]
}

func main() {
	fmt.Println(
		champagneTower(100000009, 33, 17),
	)
}
