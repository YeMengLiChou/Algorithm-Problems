package main

func maxVowels(s string, k int) int {
	res, cnt, left := 0, 0, 0
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for i := range s {
		if vowels[s[i]] {
			cnt++
		}
		left = i - k + 1
		if left < 0 {
			continue
		}
		res = max(res, cnt)
		if vowels[s[left]] {
			cnt--
		}
	}
	return res
}
