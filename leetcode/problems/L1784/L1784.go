package main

func checkOnesSegment(s string) bool {
	segments := 0
	n := len(s)
	for i := 0; i < n; {
		if s[i] == '1' {
			j := i + 1
			for j < n && s[j] == s[i] {
				j++
			}
			segments++
			i = j
		} else {
			i++
		}
	}
	return segments <= 1
}
