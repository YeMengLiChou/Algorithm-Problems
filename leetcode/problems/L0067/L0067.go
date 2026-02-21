package main

import "strings"

func addBinary(a string, b string) string {
	n, m := len(a), len(b)
	cnt := 0
	var sb strings.Builder
	for i, j := n-1, m-1; i >= 0 || j >= 0; {
		x := cnt
		cnt = 0
		if i >= 0 {
			x += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			x += int(b[j] - '0')
			j--
		}
		if x >= 2 {
			cnt = 1
			x %= 2
		}
		sb.WriteRune(rune(x) + '0')
	}
	if cnt >= 1 {
		sb.WriteRune(rune(cnt) + '0')
	}
	raw := sb.String()
	n = len(raw)
	res := make([]byte, n)
	for i, x := range raw {
		res[n-i-1] = byte(x)
	}
	return string(res)
}
