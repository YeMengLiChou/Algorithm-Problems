package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func A(n, m int) int {
	ans := 1
	end := m - n + 1
	for idx := m; idx >= end; idx-- {
		ans *= idx
	}
	return ans
}

func C(n, m int) int {
	return A(n, m) / A(n, n)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	relation := make(map[int]map[int]bool)

	fmt.Fscan(in, &n, &m)
	for range m {
		var a, b int
		fmt.Fscan(in, &a, &b)
		elem, ok := relation[a]
		if !ok {
			elem = make(map[int]bool)
			relation[a] = elem
		}
		elem[b] = true

		elem, ok = relation[b]
		if !ok {
			elem = make(map[int]bool)
			relation[b] = elem
		}
		elem[a] = true
	}

	for idx := range n {
		man := idx + 1
		elem, ok := relation[man]
		available := n - 1
		if ok {
			available = n - len(elem) - 1
		}
		fmt.Println(available)
		ans := 0
		if available >= 3 {
			ans = C(3, available)
		}
		out.WriteString(strconv.Itoa(ans))
		out.WriteString(" ")
	}
}
