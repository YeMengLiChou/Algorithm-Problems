package main

import "fmt"

type Node struct {
	Val int
}

func main() {
	nodes := make([]Node, 5)
	for i, x := range nodes {
		x.Val = i
	}
	fmt.Println(nodes)
}
