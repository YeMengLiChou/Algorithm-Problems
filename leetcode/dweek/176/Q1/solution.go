package main

import (
	"strings"
)

func mapWordWeights(words []string, weights []int) string {
	var b strings.Builder
	for _, x := range words {
		cnt := 0
		for _, c := range x {
			cnt += weights[int(c-'a')]
		}
		cnt %= 26
		b.WriteRune(rune('a' + (26 - cnt - 1)))
	}
	return b.String()
}
