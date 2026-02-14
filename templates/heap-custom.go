package templates

import (
	"cmp"
	"container/heap"
)

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

func (hp *Heap[T]) Len() int               { return len(hp.d) }
func (hp *Heap[T]) Less(i int, j int) bool { return hp.cmp(hp.d[i], hp.d[j]) }
func (hp *Heap[T]) Swap(i int, j int)      { hp.d[i], hp.d[j] = hp.d[j], hp.d[i] }
func (hp *Heap[T]) Push(x any)             { hp.d = append(hp.d, x.(T)) }
func (hp *Heap[T]) Pop() any               { n := len(hp.d); item := hp.d[n-1]; hp.d = hp.d[:n-1]; return item }
func (hp *Heap[T]) Peek() *T {
	n := hp.Len()
	if n == 0 {
		return nil
	} else {
		return &hp.d[n-1]
	}
}
