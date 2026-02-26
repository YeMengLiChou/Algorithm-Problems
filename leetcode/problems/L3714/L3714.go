package main

import "fmt"

func main() {
	fmt.Println(longestBalanced("abbac"))
}

func longestBalanced(s string) int {
	n, res := len(s), 0
	// check only one character
	for i := 0; i < n; {
		j := i
		for j < n && s[i] == s[j] {
			j++
		}
		res = max(res, j-i)
		i = j
	}

	// check two characters
	check := func(a, b byte) {
		for i := 0; i < n; {
			sum := 0
			mp := map[int]int{0: i - 1}
			for i < n && (s[i] == a || s[i] == b) {
				if s[i] == b {
					sum += -1
				} else {
					sum += 1
				}
				if pre, ok := mp[sum]; ok {
					res = max(res, i-pre)
				} else {
					mp[sum] = i
				}
				i++
			}
			i++
		}
	}

	check('a', 'b')
	check('a', 'c')
	check('b', 'c')

	// check three characters
	cnt := map[rune]int{}
	mp := map[pair]int{{}: -1}
	for i, x := range s {
		cnt[x]++
		p := pair{cnt['a'] - cnt['b'], cnt['b'] - cnt['c']}
		if pre, ok := mp[p]; ok {
			res = max(res, i-pre)
		} else {
			mp[p] = i
		}
	}
	return res
}

type pair struct{ i, j int }
