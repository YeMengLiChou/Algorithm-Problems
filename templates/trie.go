package templates

type Trie struct {
	ch    [26]*Trie
	isEnd bool
}

func (tr *Trie) Insert(word string) {
	u := tr
	for _, c := range word {
		c -= 'a'
		if u.ch[c] == nil {
			u.ch[c] = &Trie{}
		}
		u = u.ch[c]
	}
	u.isEnd = true
}

func (tr *Trie) Find(word string) bool {
	u := tr
	for _, c := range word {
		c -= 'a'
		u = u.ch[c]
		if u == nil {
			return false
		}
	}
	return u.isEnd
}
