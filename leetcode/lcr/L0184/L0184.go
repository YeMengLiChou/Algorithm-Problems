package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Add(4)
	obj.Add(7)
	fmt.Println(obj.Get_max())
	fmt.Println(obj.Remove())
	fmt.Println(obj.Get_max())
}

type Checkout struct {
	hh, tt, idx int
	dq          []int
	a           []int
}

func Constructor() Checkout {
	return Checkout{
		0, -1, 0,
		make([]int, 1<<5),
		make([]int, 0),
	}
}

func (this *Checkout) Get_max() int {
	if this.hh > this.tt {
		return -1
	}
	return this.a[this.dq[this.hh]]
}

func (this *Checkout) Add(value int) {
	this.a = append(this.a, value)
	idx := len(this.a) - 1
	for this.tt >= this.hh && this.a[this.dq[this.tt]] <= value {
		this.tt--
	}
	this.tt++
	if this.tt >= len(this.dq) {
		this.dq = append(this.dq, idx)
	} else {
		this.dq[this.tt] = idx
	}
}

func (this *Checkout) Remove() int {
	if this.idx == len(this.a) {
		return -1
	} else {
		if this.dq[this.hh] == this.idx {
			this.hh++
		}
		this.idx++
		return this.a[this.idx-1]
	}
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */
