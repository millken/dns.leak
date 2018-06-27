package main

import (
	"container/list"
)

const maxItemSize = 1000

type LRUCache struct {
	items    map[string]*list.Element // key -> element
	elements *list.List
	size     int32
}

func NewLRUCache() *LRUCache {
	return &LRUCache{items: make(map[string]*list.Element), elements: list.New()}
}

func (c *LRUCache) Put(key string, val interface{}) {
	v, ok := c.items[key]
	if ok {
		v.Value = val
		c.elements.MoveToFront(v)
		return
	}

	if c.Size() >= maxItemSize {
		c.removeOldestItem()
	}

	first := c.elements.Front()
	var e *list.Element
	if first != nil {
		e = c.elements.InsertBefore(val, first)
	} else {
		e = c.elements.PushFront(val)
	}

	c.items[key] = e
	c.size++
}

func (c *LRUCache) Get(key string) interface{} {
	if v, ok := c.items[key]; ok {
		c.elements.MoveToFront(v)
		return v.Value
	}

	return nil
}

func (c *LRUCache) Remove(key string) {
	if v, ok := c.items[key]; ok {
		c.elements.Remove(v)
		delete(c.items, key)
		c.size--
	}
}

func (c *LRUCache) removeOldestItem() {
	c.elements.Remove(c.elements.Back())
	c.size--
}

func (c *LRUCache) Size() int32 {
	return c.size
}
