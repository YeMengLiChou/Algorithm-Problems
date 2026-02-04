package main

import (
	"container/heap"
	"sort"
)

// minimumCost 宽度为 dist 的滑动窗口内选出k-1个最小值
func minimumCost(nums []int, k int, dist int) int64 {
	k -= 1
	var (
		n = len(nums)
		l = 1
		r = 1 + dist + 1
	)
	L := &LazyHeap{todo: map[int]int{}}
	R := &LazyHeap{todo: map[int]int{}}

	for _, x := range nums[1:r] {
		L.push(x)
	}
	for L.size > k {
		R.push(-L.pop())
	}
	res := L.sum
	for r < n {
		remove, add := nums[l], nums[r]
		// 移除
		if remove <= L.peek() {
			L.del(remove)
		} else {
			R.del(-remove)
		}
		// 添加
		if add < L.peek() {
			L.push(add)
		} else {
			R.push(-add)
		}

		if L.size < k {
			L.push(-R.pop())
		} else if L.size > k {
			R.push(-L.pop())
		}
		res = min(res, L.sum)
		l++
		r++
	}
	return res + int64(nums[0])
}

type LazyHeap struct {
	sort.IntSlice
	todo map[int]int
	size int
	sum  int64
}

// Less 当需要改为大根堆时，才重新实现 Less
func (lh *LazyHeap) Less(i, j int) bool { return lh.IntSlice[i] > lh.IntSlice[j] }
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
	lh.sum -= int64(x)
}

func (lh *LazyHeap) push(x int) {
	if lh.todo[x] > 0 {
		lh.todo[x]--
	} else {
		heap.Push(lh, x)
	}
	lh.size++
	lh.sum += int64(x)
}

func (lh *LazyHeap) pop() int {
	lh.do()
	lh.size--
	v := heap.Pop(lh).(int)
	lh.sum -= int64(v)
	return v
}

func (lh *LazyHeap) peek() int { lh.do(); return lh.IntSlice[0] }

func (lh *LazyHeap) do() {
	for lh.Len() > 0 && lh.todo[lh.IntSlice[0]] > 0 {
		lh.todo[lh.IntSlice[0]]--
		heap.Pop(lh)
	}
}
