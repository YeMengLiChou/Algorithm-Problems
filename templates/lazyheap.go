package templates

import (
	"container/heap"
	"sort"
)

type LazyHeap struct {
	sort.IntSlice
	todo map[int]int
	size int
}

// Less 当需要改为大根堆时，才重新实现 Less
func (lh *LazyHeap) Less(i, j int) bool { return lh.IntSlice[i] < lh.IntSlice[j] }
func (lh *LazyHeap) Push(x any)         { lh.IntSlice = append(lh.IntSlice, x.(int)) }
func (lh *LazyHeap) Pop() any {
	n := len(lh.IntSlice)
	item := lh.IntSlice[n-1]
	lh.IntSlice = lh.IntSlice[:n-1]
	return item
}

// del 懒删除
func (lh *LazyHeap) del(x int) {
	lh.todo[x]++
	lh.size--
}

func (lh *LazyHeap) push(x int) {
	if lh.todo[x] > 0 {
		lh.todo[x]--
	} else {
		heap.Push(lh, x)
	}
	lh.size++
}

func (lh *LazyHeap) pop() int {
	lh.do()
	lh.size--
	v := heap.Pop(lh).(int)
	return v
}

func (lh *LazyHeap) peek() int { lh.do(); return lh.IntSlice[0] }

func (lh *LazyHeap) do() {
	for lh.Len() > 0 && lh.todo[lh.IntSlice[0]] > 0 {
		lh.todo[lh.IntSlice[0]]--
		heap.Pop(lh)
	}
}
