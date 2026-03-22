package main

const M = 1e9 + 7

type Fancy struct {
	data     []int
	add, mul int
}

func Constructor() Fancy {
	return Fancy{mul: 1}
}

func (f *Fancy) Append(val int) {
	f.data = append(f.data, (val-f.add+M)*qpow(f.mul, M-2, M)%M)
}

func (f *Fancy) AddAll(inc int) {
	f.add = (f.add + inc) % M
}

func (f *Fancy) MultAll(m int) {
	f.mul = f.mul * m % M
	f.add = f.add * m % M
}

func (f *Fancy) GetIndex(idx int) int {
	if idx >= len(f.data) {
		return -1
	}
	return (f.data[idx]*f.mul + f.add) % M
}

func qpow(x, y, m int) int {
	ans := 1
	for ; y > 0; y >>= 1 {
		if y&1 == 1 {
			ans = ans * x % m
		}
		x = x * x % m
	}
	return ans
}
