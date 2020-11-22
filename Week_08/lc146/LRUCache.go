package main

import (
	"container/list"
)

type KV struct {
	k  	int
	v	int
}

type LRUCache struct {
	cap 	int
	table   map[int] *list.Element
	l		*list.List
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{
		capacity,
		make(map[int] *list.Element),
		list.New(),
	}
	return ret
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.table[key]
	if !ok {
		return -1
	}
	this.l.MoveToFront(e)
	kv, _ := e.Value.(KV)
	return kv.v
}

func (this *LRUCache) Put(key int, value int) {
	//remove old if exist
	e, ok := this.table[key]
	if ok {
		this.Remove(e)
	}
	//remove back if full
	if this.l.Len() == this.cap {
		this.Remove(this.l.Back())
	}
	this.l.PushFront(KV{key,value})
	this.table[key] = this.l.Front()
}

func (this *LRUCache) Remove(e *list.Element) {
	if e == nil {return}
	kv,_ := e.Value.(KV)
	this.l.Remove(e)
	delete(this.table, kv.k)
}