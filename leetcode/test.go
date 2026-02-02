package main

import (
	"fmt"
	"slices"
)

func main() {
	l := []int{1, 2, 2, 2, 4, 4, 5}
	x, ok := slices.BinarySearch(l, 3)
	fmt.Println(l[x], ok)
}
