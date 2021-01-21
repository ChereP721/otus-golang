package hw04_lru_cache //nolint:golint,stylecheck
import (
	"fmt"
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*cacheItem
}

type cacheItem struct {
	value     interface{}
	queueItem *listItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEl, ok := c.items[key]
	if ok {
		cacheEl.value = value
		c.queue.MoveToFront(cacheEl.queueItem)

		return true
	}

	if len(c.items) == c.capacity {
		c.pushOut()
	}

	newQueueItem := c.queue.PushFront(key)
	c.items[key] = &cacheItem{
		value:     value,
		queueItem: newQueueItem,
	}

	return false
}

func (c *lruCache) pushOut() {
	delKey, ok := c.queue.Back().Value.(Key)
	if !ok {
		// поскольку в интерфейсе нет возврата ошибки - паникуем
		panic(fmt.Sprintf("Problem in lruCache.Set realization - wrong queue value type: expected - Key, actual - %T", c.queue.Back().Value))
	}
	c.queue.Remove(c.queue.Back())
	delete(c.items, delKey)
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEl, ok := c.items[key]
	if !ok {
		return nil, false
	}

	c.queue.MoveToFront(cacheEl.queueItem)

	return cacheEl.value, true
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*cacheItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		mu:       sync.Mutex{},
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*cacheItem, capacity),
	}
}
