package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hasAllCodes("00110110", 2))
	fmt.Println(hasAllCodes("0110", 2))
	fmt.Println(hasAllCodes("00000000001011100", 3))
}

func hasAllCodes(s string, k int) bool {
	vis := map[int]bool{}
	n := len(s)
	if k > n {
		return false
	}
	mask := (1 << k) - 1
	mx, a := int(math.Pow(2, float64(k))), 0
	for i := range n {
		a = ((a << 1) | int(s[i]-'0')) & mask
		if i < k-1 {
			continue
		}
		vis[a] = true
		if len(vis) == mx {
			return true
		}
	}
	return false
}
