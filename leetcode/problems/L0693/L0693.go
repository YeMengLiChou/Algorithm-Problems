package main

func hasAlternatingBits(n int) bool {
	pre := -1
	for n > 0 {
		if pre == n&1 {
			return false
		}
		pre = n & 1
		n >>= 1
	}
	return true
}
