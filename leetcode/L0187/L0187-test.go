package main

func findRepeatedDnaSequences(s string) []string {
	const B = 13131313
	n := len(s)
	sh := NewStringHasher(&s, B)
	mp := make(map[uint64]int)
	res := make([]string, 0)
	for i := 0; i+10 <= n; i++ {
		j := i + 10
		hash := sh.SubHash(i, j)
		cnt := mp[hash]
		if cnt == 1 {
			res = append(res, s[i:j])
		}
		mp[hash] = cnt + 1
	}
	return res
}

// StringHasher 字符串哈希
type StringHasher struct {
	h []uint64 // 存储每一位对应的hash值
	p []uint64 // 存储 p^i 值，便于求子串 hash
	s *string
}

func NewStringHasher(s *string, base uint64) *StringHasher {
	n := len(*s)
	p := make([]uint64, n+1)
	h := make([]uint64, n+1)
	p[0] = 1
	for i, c := range *s {
		p[i+1] = p[i]*base + uint64(c-'a')
		h[i+1] = h[i] * base
	}
	return &StringHasher{h, p, s}
}

// SubHash 子串 [l, r) 的哈希值
func (sh *StringHasher) SubHash(l, r int) uint64 {
	return sh.h[r] - sh.h[l]*sh.p[r-l]
}
