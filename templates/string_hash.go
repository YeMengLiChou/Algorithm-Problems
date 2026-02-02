package templates

// StringHasher 字符串哈希
type StringHasher struct {
	h []uint64 // 存储每一位对应的hash值
	p []uint64 // 存储 p^i 值，便于求子串 hash
	s *string
}

func NewStringHasher(s *string, base uint64) *StringHasher {
	n := len(*s)
	h := make([]uint64, n+1)
	p := make([]uint64, n+1)
	p[0] = 1

	for i, c := range *s {
		h[i+1] = h[i]*base + uint64(c-'a')
		p[i+1] = p[i] * base
	}

	return &StringHasher{h, p, s}
}

// SubHash 子串 [l, r) 的哈希值
func (sh *StringHasher) SubHash(l, r int) uint64 {
	// h[r-1] - h[l-1] 需要向后偏移1位
	return sh.h[r] - sh.h[l]*sh.p[r-l]
}
