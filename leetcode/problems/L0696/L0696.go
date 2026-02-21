package main

func countBinarySubstrings(s string) int {
	n := len(s)
	ans := 0
	pre, cnt := byte(2), 0
	for i := 0; i < n; {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		if pre != s[i] {
			ans += min(cnt, j-i)
			cnt = j - i
			pre = s[i]
		}
	}
	return ans
}
