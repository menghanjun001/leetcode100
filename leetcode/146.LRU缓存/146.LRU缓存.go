package leetcode

import (
	"container/list"
)

//数据结构有点问题 没ac
type LRUCache struct {
	cap int
	l   *list.List
	m   map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	//check if it exists
	if node, ok := this.m[key]; ok {
		//if it exists,get it and move it to front
		this.l.MoveToFront(node)
		return node.Value.(int)
	}
	//if not,return -1
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//check if it exists
	//if it exists,update map and move it to front
	if node, ok := this.m[key]; ok {
		//if it exists,get it and move it to front
		this.l.MoveToFront(node)
		node.Value = value
	} else {
		//if not,check if list is full
		if this.l.Len() == this.cap {
			//if full,delete tail node and refresh map
			this.l.Remove(this.l.Back())
			delete(this.m, key)
		}
		//move new node to front,and add it to map
		newNode := &list.Element{Value: value}
		this.l.PushFront(newNode)
		this.m[key] = newNode
	}
}
