package main

import "math"

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	const INF = math.MaxInt64 >> 1
	d := [26][26]int{} // 已知最多有26，直接用数组，替换 slices
	for i := range 26 {
		for j := range 26 {
			if i != j {
				d[i][j] = INF
			}
		}
	}
	// 更新已有边信息
	for i, from := range original {
		to, x := changed[i], cost[i]
		d[from-'a'][to-'a'] = min(d[from-'a'][to-'a'], x)
	}

	for k := range 26 {
		for i := range 26 {
			if d[i][k] == INF { // 边不连通时，直接跳过
				continue
			}
			for j := range 26 {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	var ans int64
	for i := range source {
		from, to := source[i]-'a', target[i]-'a'
		if from == to {
			continue
		}
		c := d[from][to]
		if c >= INF {
			return -1
		}
		ans += int64(c)
	}
	return ans
}
