package main

import (
	"strconv"
	"strings"
)

func parse(q *string) (int, int, int) {
	s := strings.Split(*q, " ")
	action := 0
	if s[0] == "queries" {
		action = 1
	}
	a, _ := strconv.Atoi(s[1])
	b, _ := strconv.Atoi(s[1])
	return action, a, b
}

func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
}
