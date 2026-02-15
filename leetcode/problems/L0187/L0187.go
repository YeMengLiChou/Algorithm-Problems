package main

func findRepeatedDnaSequences(s string) []string {
	const B = 13131313
	n := len(s)
	h := make([]uint64, n+1)
	p := make([]uint64, n+1)
	p[0] = 1
	for i, c := range s {
		h[i+1] = h[i]*B + uint64(c-'a')
		p[i+1] = p[i] * B
	}
	mp := make(map[uint64]int)
	res := make([]string, 0)
	for i := 1; i+9 <= n; i++ {
		j := i + 9 //
		hash := h[j] - h[i-1]*p[j-i+1]
		cnt := mp[hash]
		if cnt == 1 {
			res = append(res, s[i-1:j])
		}
		mp[hash] = cnt + 1
		j++
	}
	return res
}
