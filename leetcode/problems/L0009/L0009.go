package main

import "strconv"

func isPalindromeV1(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for l, r := 0, len(s)-1; l < r; {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}
}

func isPalindrome(x int) bool {
	return isPalindromeV2(x)
}
