package main

import "fmt"

func reverseBitsV1(n int) int {
	ans := 0
	idx := 31
	for idx >= 0 {
		ans |= (n & 1) << idx
		n >>= 1
		idx--
	}
	return ans
}

var (
	m1 = 0b01010101010101010101010101010101
	m2 = 0b00110011001100110011001100110011
	m3 = 0b00001111000011110000111100001111
	m4 = 0b00000000111111110000000011111111
	m5 = 0b00000000000000001111111111111111
)

func reverseBitsV2(n int) int {
	n = (n >> 1 & m1) | (n & m1 << 1)
	n = (n >> 2 & m2) | (n & m2 << 2)
	n = (n >> 4 & m3) | (n & m3 << 4)
	n = (n >> 8 & m4) | (n & m4 << 8)
	n = (n >> 16 & m5) | (n & m5 << 16)
	return n
}

func reverseBits(n int) int {
	return reverseBitsV2(n)
}

func main() {
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(2147483644))
}
