package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isDigitorialPermutation(10))
}

func isDigitorialPermutation(n int) bool {
	s := strconv.Itoa(n)
	sum := int64(0)
	mp := map[int]int{}
	for _, x := range s {
		cnt := int64(1)
		digital := int(x - '0')
		mp[digital]++
		for digital > 0 {
			cnt *= int64(digital)
			digital--
		}
		sum += cnt
	}
	return dfs(0, len(s), 0, sum, mp)
}

func dfs(idx int, length int, cnt int64, sum int64, digitals map[int]int) bool {
	if idx == length {
		return cnt == sum
	}
	start := 0
	if cnt == 0 {
		start = 1
	}
	for i := start; i <= 9; i++ {
		if digitals[i] > 0 {
			digitals[i]--
			if dfs(idx+1, length, cnt*10+int64(i), sum, digitals) {
				return true
			}
			digitals[i]++
		}
	}
	return false
}
