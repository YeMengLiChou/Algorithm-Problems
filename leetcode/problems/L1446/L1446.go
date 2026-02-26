package main

func maxPower(s string) int {
	n, res := len(s), 0
	for i := 0; i < n; {
		j := i + 1
		for j < n && s[i] == s[j] {
			j++
		}
		res = max(res, j-i)
		i = j
	}
	return res
}
