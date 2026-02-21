package main

// import "slices"

// type Trie struct {
// 	next  [][26]int32
// 	isEnd []bool
// 	idx   int32
// }

// func Constructor() Trie {
// 	return Trie{
// 		make([][26]int32, 1),
// 		make([]bool, 1),
// 		0,
// 	}
// }

// func (trie *Trie) Insert(word string) {
// 	var u int32 = 0
// 	for _, c := range word {
// 		t := c - 'a'
// 		if trie.next[u][t] == 0 {
// 			trie.idx++
// 			// 动态扩大数组长度，绕开一次申请大内存
// 			// 同时扩大的长度应该偏大，避免频繁触发系统调用带来的开销.
// 			if cap(trie.next) <= int(trie.idx) {
// 				trie.next = slices.Grow(trie.next, 1000)
// 				trie.isEnd = slices.Grow(trie.isEnd, 1000)
// 			}
// 			trie.next = append(trie.next, [26]int32{})
// 			trie.isEnd = append(trie.isEnd, false)
// 			trie.next[u][t] = trie.idx
// 		}
// 		u = trie.next[u][t]
// 	}
// 	trie.isEnd[u] = true
// }

// func (trie *Trie) Search(word string) bool {
// 	var u int32 = 0
// 	for _, c := range word {
// 		t := c - 'a'
// 		if trie.next[u][t] == 0 {
// 			return false
// 		}
// 		u = trie.next[u][t]
// 	}
// 	return trie.isEnd[u]
// }

// func (trie *Trie) StartsWith(prefix string) bool {
// 	var u int32 = 0
// 	for _, c := range prefix {
// 		t := c - 'a'
// 		if trie.next[u][t] == 0 {
// 			return false
// 		}
// 		u = trie.next[u][t]
// 	}
// 	return true
// }

// /**
//  * Your Trie object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Insert(word);
//  * param_2 := obj.Search(word);
//  * param_3 := obj.StartsWith(prefix);
//  */
