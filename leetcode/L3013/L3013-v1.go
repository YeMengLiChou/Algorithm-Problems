package main

import (
	"cmp"
	"container/heap"
	"slices"
)

// minimumCost 宽度为 dist 的滑动窗口内选出k-1个最小值
func minimumCost(nums []int, k int, dist int) int64 {
	k -= 1
	var (
		n = len(nums)
		l = 1
		r = 1 + dist + 1
	)

	lh := NewMaxHeapOrdered(append(make([]int, 0), nums[1:r]...))
	rh := NewMinHeapOrdered(make([]int, 0))
	sum := int64(nums[0]) + Sum(lh)

	L2R := func() {
		x := heap.Pop(lh).(int)
		heap.Push(rh, x)
		sum -= int64(x)
	}
	R2L := func() {
		x := heap.Pop(rh).(int)
		heap.Push(lh, x)
		sum += int64(x)
	}
	for lh.Len() > k {
		L2R()
	}
	res := sum
	for r < n {
		remove, add := nums[l], nums[r]
		// 移除离开窗口的元素
		idx := slices.Index(lh.d, remove)
		if idx >= 0 {
			heap.Remove(lh, idx)
			sum -= int64(remove)
		} else {
			idx = slices.Index(rh.d, remove)
			heap.Remove(rh, idx)
		}

		// 添加新元素
		if add < (*lh.Peek()) {
			heap.Push(lh, add)
			sum += int64(add)
		} else {
			heap.Push(rh, add)
		}

		if lh.Len() < k {
			R2L()
		} else if lh.Len() > k {
			L2R()
		}

		l++
		r++
		res = min(res, sum)
	}
	return res
}

func Sum(hp *Heap[int]) int64 {
	res := int64(0)
	for _, x := range hp.d {
		res += int64(x)
	}
	return res
}

type Heap[T any] struct {
	d   []T
	cmp func(a, b T) bool
}

func NewHeap[T any](d []T, cmp func(a, b T) bool) *Heap[T] {
	hp := &Heap[T]{d, cmp}
	heap.Init(hp)
	return hp
}

// NewMinHeapOrdered 小根堆
func NewMinHeapOrdered[T cmp.Ordered](d []T) *Heap[T] {
	return NewHeap(d, func(a, b T) bool { return a < b })
}

// NewMaxHeapOrdered 大根堆
func NewMaxHeapOrdered[T cmp.Ordered](d []T) *Heap[T] {
	return NewHeap(d, func(a, b T) bool { return b < a })
}

func (hp *Heap[T]) Peek() *T {
	n := hp.Len()
	if n == 0 {
		return nil
	} else {
		return &hp.d[0]
	}
}
func (hp *Heap[T]) Len() int               { return len(hp.d) }
func (hp *Heap[T]) Less(i int, j int) bool { return hp.cmp(hp.d[i], hp.d[j]) }
func (hp *Heap[T]) Swap(i int, j int)      { hp.d[i], hp.d[j] = hp.d[j], hp.d[i] }
func (hp *Heap[T]) Push(x any)             { hp.d = append(hp.d, x.(T)) }
func (hp *Heap[T]) Pop() any               { n := len(hp.d); item := hp.d[n-1]; hp.d = hp.d[:n-1]; return item }
