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
	mask, a := 0, 0
	for idx := range k {
		mask |= 1 << idx
		a <<= 1
		a |= int(s[idx] - '0')
	}
	vis[a] = true
	mx := int(math.Pow(2, float64(k)))
	for i := k; i < n; i++ {
		a <<= 1
		a |= int(s[i] - '0')
		a &= mask
		vis[a] = true
		if len(vis) == mx {
			return true
		}
	}
	return false
}

