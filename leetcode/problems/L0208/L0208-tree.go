package main

// Trie 树形式的字典树实现
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (trie *Trie) Insert(word string) {
	u := trie
	for _, c := range word {
		t := c - 'a'
		ne := u.children[t]
		if ne == nil {
			ne := &Trie{}
			u.children[t] = ne
		}
		u = ne
	}
	u.isEnd = true
}

func (trie *Trie) Search(word string) bool {
	u := trie
	for _, c := range word {
		t := c - 'a'
		ne := u.children[t]
		if ne == nil {
			return false
		}
		u = ne
	}
	return u.isEnd
}

func (trie *Trie) StartsWith(prefix string) bool {
	u := trie
	for _, c := range prefix {
		t := c - 'a'
		ne := u.children[t]
		if ne == nil {
			return false
		}
		u = ne
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
