package main

type LRUCache struct {
	cache      map[int]*DLinkedNode
	size       int
	capacity   int
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	Prev, Next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		capacity: capacity,
		size:     0,
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size += 1
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size -= 1
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}

}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.Prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.Prev = this.head
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}
