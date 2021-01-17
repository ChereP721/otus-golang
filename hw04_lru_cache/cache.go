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
	items    map[Key]*listItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func getCacheItemPtr(value interface{}) *cacheItem {
	ciPtr, ok := value.(*cacheItem)
	if !ok {
		// поскольку в интерфейсе нет возврата ошибки - паникуем
		panic("Wrong cache item type!")
	}

	return ciPtr
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	li, ok := cache.items[key]
	if ok {
		getCacheItemPtr(li.Value).value = value
		cache.queue.MoveToFront(li)

		return true
	}

	if len(cache.items) == cache.capacity {
		delKey := getCacheItemPtr(cache.queue.Back().Value).key
		cache.queue.Remove(cache.queue.Back())
		delete(cache.items, delKey)
	}

	newItem := cache.queue.PushFront(&cacheItem{
		key:   key,
		value: value,
	})
	cache.items[key] = newItem

	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	li, ok := cache.items[key]
	if !ok {
		return nil, false
	}

	cache.queue.MoveToFront(li)
	return getCacheItemPtr(li.Value).value, true
}

func (cache *lruCache) Clear() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	for k, li := range cache.items {
		cache.queue.Remove(li)
		delete(cache.items, k)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		mu:       sync.Mutex{},
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*listItem),
	}
}
