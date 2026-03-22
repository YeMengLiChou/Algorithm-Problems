package main

func bitwiseComplement(n int) int {
	idx := 32
	for idx >= 0 && (n>>idx)&1 == 0 {
		idx--
	}
	if idx < 0 {
		return 1
	}
	res := 0
	for idx >= 0 {
		res |= (1 - (n >> idx & 1)) << idx
		idx--
	}
	return res
}
