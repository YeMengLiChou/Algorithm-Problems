package main

import "container/heap"

type pair struct {
	sum, idx int
}

type hp []pair

func (h hp) Len() int { return len(hp) }
func (h hp) Less(i int, j int) bool {
	if h[i].sum == h[j].sum {
		return h[i].idx < h[j].idx
	}
	return h[i].sum < h[j].sum
}

func (h hp) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func minimumPairRemoveal(nums []int) int {
	size, desc := len(nums), 0
	values := make(hp, size)
	for i := range len(nums) - 1 {
		if nums[i] > nums[i+1] {
			desc++
		}
		values[i] = pair{nums[i] + nums[i+1], i}
	}
	heap.Init(h, values)
}
