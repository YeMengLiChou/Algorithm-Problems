package main

import "fmt"

func main() {
	fmt.Println(
		numberOfStableArrays(24, 35, 21),
	)
}

const M int = 1e9 + 7

func numberOfStableArrays(zero int, one int, limit int) int {
	memo := make([][][2]int, zero+1)
	for i := range memo {
		memo[i] = make([][2]int, one+1)
		for j := range memo[i] {
			memo[i][j] = [...]int{-1, -1}
		}
	}
	return (dfs(zero, one, 1, limit, memo) + dfs(zero, one, 0, limit, memo)) % M
}

func dfs(zero, one, k, limit int, memo [][][2]int) int {
	p := &memo[zero][one][k]
	if *p != -1 {
		return *p
	}
	v := 0
	if zero <= 0 {
		if k == 1 && one <= limit {
			// 没有 0 时只能填 1，且需要满足 limit，只有 1 种方案
			v = 1
		}
	} else if one <= 0 {
		if k == 0 && zero <= limit {
			// 没有 1 时只能填 0，且需要满足 limi ，只有 1 种方案
			v = 1
		}
	} else if k == 0 {
		v = (dfs(zero-1, one, 0, limit, memo) + dfs(zero-1, one, 1, limit, memo)) % M
		if zero > limit {
			v = (v - dfs(zero-limit-1, one, 1, limit, memo) + M) % M
		}
	} else {
		v = (dfs(zero, one-1, 0, limit, memo) + dfs(zero, one-1, 1, limit, memo)) % M
		if one > limit {
			v = (v - dfs(zero, one-limit-1, 0, limit, memo) + M) % M
		}
	}
	*p = v
	return v
}
