package main

func prefixConnected(words []string, k int) int {
	res := 0
	g := map[string]int{}
	for _, x := range words {
		if len(x) < k {
			continue
		}
		prefix := x[:k]
		g[prefix]++
	}
	for _, v := range g {
		if v >= 2 {
			res++
		}
	}
	return res
}
