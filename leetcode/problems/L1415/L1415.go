package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(getHappyStringV1(1, 3))
	fmt.Println(getHappyStringV1(3, 9))
}

func getHappyString(n int, k int) string {
	res := make([]string, 0)
	builder := make([]byte, n)
	build(&res, &builder, 0, 0, n)
	if len(res) < k {
		return ""
	}
	slices.Sort(res)
	return res[k-1]
}

func build(result *[]string, sb *[]byte, pre byte, idx, n int) {
	if idx == n {
		*result = append(*result, string(*sb))
		return
	}

	for i := range byte(3) {
		x := byte('a') + i
		if x == pre {
			continue
		}
		(*sb)[idx] = x
		build(result, sb, x, idx+1, n)
	}
}

func getHappyStringV1(n int, k int) string {
	if k > 3<<(n-1) {
		return ""
	}
	k--
	sb := make([]byte, n)
	for i := range sb {
		sb[i] = 'a'
	}

	sb[0] += byte(k >> (n - 1))

	for i := 1; i < n; i++ {
		v := 'a' + byte(k>>(n-1-i)&1)
		if v >= sb[i-1] {
			sb[i] = v + 1
		} else {
			sb[i] = v
		}
	}
	return string(sb)
}
