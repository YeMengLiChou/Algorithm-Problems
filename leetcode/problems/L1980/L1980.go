package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findDifferentBinaryStringV1(nums []string) string {
	n := len(nums)
	vis := make(map[int]bool, n)
	for _, s := range nums {
		x, _ := strconv.ParseInt(s, 2, 32)
		vis[int(x)] = true
	}
	for i := 0; ; i++ {
		if !vis[i] {
			return fmt.Sprintf("%0"+strconv.Itoa(n)+"b", i)
		}
	}
}

func findDifferentBinaryStringV2(nums []string) string {
	var sb strings.Builder
	n := len(nums)
	for i := range n {
		x := nums[i][i]
		sb.WriteByte('1' - x + '0')
	}
	return sb.String()
}

func findDifferentBinaryString(nums []string) string {
	return findDifferentBinaryStringV2(nums)
}
