package main

import "slices"

func nextGreatestLetter(letters []byte, target byte) byte {
	pos, _ := slices.BinarySearchFunc(letters, target, func(c, t byte) int {
		if c > t {
			return 0
		} else {
			return -1
		}
	})
	if pos >= len(letters) {
		return letters[0]
	}
	return letters[pos]
}
