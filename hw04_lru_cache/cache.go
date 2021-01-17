package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

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

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cacheEl, ok := cache.items[key]
	if ok {
		cacheEl.value = value
		cache.queue.MoveToFront(cacheEl.queueItem)

		return true
	}

	if len(cache.items) == cache.capacity {
		delKey, ok := cache.queue.Back().Value.(Key)
		if !ok {
			// поскольку в интерфейсе нет возврата ошибки - паникуем
			panic("Wrong cache item type!")
		}
		cache.queue.Remove(cache.queue.Back())
		delete(cache.items, delKey)
	}

	newQueueItem := cache.queue.PushFront(key)
	cache.items[key] = &cacheItem{
		value:     value,
		queueItem: newQueueItem,
	}

	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cacheEl, ok := cache.items[key]
	if !ok {
		return nil, false
	}

	cache.queue.MoveToFront(cacheEl.queueItem)

	return cacheEl.value, true
}

func (cache *lruCache) Clear() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	for key, cacheEl := range cache.items {
		cache.queue.Remove(cacheEl.queueItem)
		delete(cache.items, key)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		mu:       sync.Mutex{},
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*cacheItem),
	}
}
