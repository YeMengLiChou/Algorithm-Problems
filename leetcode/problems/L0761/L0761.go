package main

import (
	"slices"
	"strings"
)

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	substrings := []string{}
	diff := 0
	start := 0
	for i, x := range s {
		if x == '1' {
			diff++
		} else {
			diff--

			if diff == 0 {
				substrings = append(substrings, makeLargestSpecial(s[start:i+1]))
				start = i + 1
			}
		}
	}
	slices.Sort(substrings)
	slices.Reverse(substrings)
	return strings.Join(substrings, "")
}
