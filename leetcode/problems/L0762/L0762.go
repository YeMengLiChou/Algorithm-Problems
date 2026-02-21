package main

import (
	"math/bits"
	"slices"
)

func countBits(x int) int {
	cnt := 0
	for x > 0 {
		if x&1 == 1 {
			cnt++
		}
		x >>= 1
	}
	return cnt
}

func countPrimeSetBits(left int, right int) int {
	primes := map[int]bool{
		2: true, 3: true, 5: true, 7: true, 11: true,
		13: true, 17: true, 19: true, 23: true, 29: true, 31: true,
	}
	cnt := 0
	for left <= right {
		if primes[countBits(left)] {
			cnt++
		}
		left++
	}
	return cnt
}

func countPrimeSetBitsOpt(left int, right int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	cnt := 0
	for left <= right {
		if slices.Contains(primes, bits.OnesCount(uint(left))) {
			cnt++
		}
		left++
	}
	return cnt
}
