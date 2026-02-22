package main

import "strings"

func maximumXor(s string, t string) string {
	sb := strings.Builder{}
	mp := [2]int{}
	for _, x := range t {
		mp[x-'0']++
	}

	for _, x := range s {
		a := x - '0'
		b := 1 - a
		if mp[b] > 0 {
			sb.WriteRune('1')
			mp[b]--
		} else {
			mp[a]--
			sb.WriteRune('0')
		}
	}
	return sb.String()
}
