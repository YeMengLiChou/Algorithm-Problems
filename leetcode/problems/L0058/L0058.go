package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

func lengthOfLastWord(s string) int {
	idx := len(s) - 1
	for idx >= 0 && s[idx] == ' ' {
		idx--
	}
	d := idx
	for idx >= 0 && s[idx] != ' ' {
		idx--
	}
	return d - idx
}
