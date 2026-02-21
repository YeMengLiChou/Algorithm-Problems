package main

func lengthOfLongestSubstring(s string) int {
	res := 0
	mp := make(map[byte]int)
	l, r, n := 0, 0, len(s)

	for r < n {
		mp[s[r]] += 1
		for l < r && mp[s[r]] > 1 {
			mp[s[l]]--
			l++
		}
		res = max(res, r-l-1)
		r++
	}
	return res
}
