package main

func minFlips(s string) int {
	n := len(s)
	res := 1 << 30
	cnt := 0
	left := 0
	for i := range (n << 1) - 1 {
		if int(s[i%n])&1 != i&1 {
			cnt++
		}
		left = i - n + 1
		if left < 0 {
			continue
		}
		res = min(res, cnt, n-cnt)
		if int(s[left%n])&1 != left&1 {
			cnt--
		}
	}
	return res
}
