package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSteps("1101")) // 6
	fmt.Println(numSteps("10"))   // 1
	fmt.Println(numSteps("11"))   // 1
}

func numSteps(s string) int {
	res, idx := 0, len(s)-1
	cnt := 0
	for idx > 0 {
		v := int(s[idx]-'0') + cnt
		cnt = max(0, cnt-1)
		if v&1 == 1 {
			cnt += 1
			res += 2
		} else {
			res += 1
			cnt += v / 2
		}
		idx--
	}
	if cnt > 0 {
		res += 1
	}
	return res
}

