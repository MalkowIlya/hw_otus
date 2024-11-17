package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.items[key]; ok {
		c.items[key].Value = value
		c.queue.MoveToFront(c.items[key])

		c.queue.Front().Key = key
		c.items[key] = c.queue.Front()
		return true
	}

	if c.queue.Len() == c.capacity {
		delete(c.items, c.queue.Back().Key)
		c.queue.Remove(c.queue.Back())
	}

	c.queue.PushFront(value)
	c.queue.Front().Key = key
	c.items[key] = c.queue.Front()

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(c.items[key])
		c.items[key] = c.queue.Front()
		return item.Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
