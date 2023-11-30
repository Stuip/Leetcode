package main

import "fmt"

const SIZE = 26

type Trie struct {
	children [SIZE]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}


func (t *Trie) insert(word string) {
	cur := t
	n := len(word)
	for i:=0;i<n;i++{
		index := word[i] - 'a'
		if cur.children[index] == nil {
			cur.children[index] = &Trie{}
		}
		cur = cur.children[index]
	}
	cur.isEnd = true
}

func (t *Trie) Search(word string) bool {
	var res *Trie = t.find(word)
	return res != nil && res.isEnd
}


func (t *Trie) find(word string) *Trie {
	cur, n := t, len(word)
	for i:=0;i<n;i++{
		index := word[i] - 'a'
		if cur.children[index] != nil {
			cur = cur.children[index]
		} else {
			return nil
		}
	}
	return cur
}

func (t *Trie) StartsWith(prefix string) bool {
	var res *Trie = t.find(prefix)
	return res != nil
}



func main() {
	t := Constructor()
	t.insert("hello")
	fmt.Println(t.Search("hello"))
}
