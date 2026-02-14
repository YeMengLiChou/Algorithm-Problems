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

var p []int

func Swap(i, j int)    { p[i], p[j] = p[j], p[i] }
func Parent(x int) int { return (x - 1) >> 1 }
func RChild(x int) int { return (x << 1) + 2 }
func LChild(x int) int { return (x << 1) + 1 }

func Up(i int) {
	for x := Parent(i); i > 0 && p[i] > p[x]; x = Parent(i) {
		Swap(i, x)
		i = x
	}
}

// func Down(i int) {
// 	n := len(p)
// 	for
// }
