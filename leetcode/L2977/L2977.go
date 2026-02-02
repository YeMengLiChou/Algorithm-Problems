package main

import (
	"math"
)

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	const INF = math.MaxInt >> 1
	var (
		idx  int
		n    = len(original)
		root = &Trie{idx: -1}
	)
	m := n << 1
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, m)
		for j := range dis[i] {
			dis[i][j] = INF
		}
		dis[i][i] = 0
	}
	for i := range n {
		x, y, c := root.Insert(&original[i], &idx), root.Insert(&changed[i], &idx), cost[i]
		dis[x][y] = min(dis[x][y], c)
	}
	// floyd
	for k := range m {
		for i := range m {
			if dis[i][k] == INF {
				continue
			}
			for j := range m {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	m = len(source)
	f := make([]int64, m)
	for i := 1; i < m; i++ {
		f[i] = INF
	}
	for i := range m {
		var base int64 = 0
		if i > 0 {
			base = f[i-1]
		}
		if source[i] == target[i] {
			f[i] = min(f[i], base)
		}
		// 查找是否存在对应的子串
		u, v := root, root
		for j := i; j < m; j++ {
			u, v = u.son[source[j]-'a'], v.son[target[j]-'a']
			// 没有该子串
			if u == nil || v == nil {
				break
			}
			// 某个子串的前缀，还没结束
			if u.idx == -1 || v.idx == -1 || dis[u.idx][v.idx] == INF {
				continue
			}
			f[j] = min(f[j], base+int64(dis[u.idx][v.idx]))
		}
	}
	res := f[m-1]
	if res == INF {
		return -1
	}
	return res
}

type Trie struct {
	son [26]*Trie
	idx int
}

func (tr *Trie) Insert(str *string, idx *int) int {
	u := tr
	for _, c := range *str {
		c -= 'a'
		ne := &u.son[c]
		if (*ne) == nil {
			(*ne) = &Trie{idx: -1}
		}
		u = *ne
	}
	if u.idx != -1 {
		return u.idx
	} else {
		u.idx = (*idx)
		(*idx)++
		return u.idx
	}
}
