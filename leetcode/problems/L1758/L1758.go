package main

func minOperations(s string) int {
	n := len(s)
	check := func(start byte) int {
		cnt := 0
		if s[0]-'0' != start {
			cnt++
		}
		pre := start
		for i := 1; i < n; i++ {
			x := s[i] - '0'
			if x == pre {
				cnt++
				pre = 1 - x
			} else {
				pre = x
			}
		}
		return cnt
	}
	return min(check(0), check(1))
}
