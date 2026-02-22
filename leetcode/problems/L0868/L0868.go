package main

func binaryGap(n int) int {
	pre, idx, res := -1, 1, 0
	for n > 0 {
		if n&1 == 1 {
			if pre >= 1 {
				res = max(res, idx-pre)
			}
			pre = idx
		}
		idx++
		n >>= 1
	}
	return res
}
