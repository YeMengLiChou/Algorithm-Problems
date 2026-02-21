package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(readBinaryWatch(1))
}

func lowbit(x int) int {
	return x & (-x)
}

func readBinaryWatchV1(turnedOn int) []string {
	ans := make([]string, 0)
	bits := make([]int, 60)
	for i := range 60 {
		cnt := 0
		for x := i; x > 0; {
			x -= lowbit(x)
			cnt++
		}
		bits[i] = cnt
	}

	for h := range 12 {
		cnt := bits[h]
		if cnt > turnedOn {
			continue
		}
		for m := range 60 {
			if cnt+bits[m] == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}

func readBinaryWatchV2(turnedOn int) []string {
	ans := make([]string, 0)
	for h := range uint8(12) {
		cnt := bits.OnesCount8(h)
		if cnt > turnedOn {
			continue
		}
		for m := range uint8(60) {
			if cnt+bits.OnesCount8(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}

func readBinaryWatch(turnedOn int) []string {
	return readBinaryWatchV1(turnedOn)
}
