package main

import "fmt"

func main() {
	fmt.Println(lowbit(6))
	// fmt.Println(concatenatedBinary(1))
	fmt.Println(concatenatedBinary(3))
	fmt.Println(concatenatedBinary(12))
}

func lowbit(x int) int {
	return x & (-x)
}

func concatenatedBinary(n int) int {
	const M int = 1e9 + 7
	res := 0
	idx := 0
	for i := 1; i <= n; i++ {
		if lowbit(i) == i {
			idx++
		}
		res = ((res << idx) | i) % M
	}
	return res
}
