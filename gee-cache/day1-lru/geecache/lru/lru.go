package lru

import (
	"container/list"
)

// Cache is a LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	// some code goes here
	cache := Cache{
		maxBytes:  maxBytes,
		nbytes:    int64(0),
		ll:        list.New(),
		cache:     map[string]*list.Element{},
		OnEvicted: onEvicted,
	}
	return &cache
}

// Add adds a value to the cache.
func (c *Cache) Add(key string, value Value) {
	// some code goes here
	if _, ok := c.cache[key]; ok {
		oldValue, _ := c.Get(key)
		c.nbytes += int64(value.Len()) - int64(oldValue.Len())
		c.cache[key] = &list.Element{Value: &entry{key: key, value: value}}
	} else {
		element := c.ll.PushFront(&entry{key: key, value: value})
		c.nbytes += int64(value.Len()) + int64(len(key))
		c.cache[key] = element
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	// some code goes here
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		kv := element.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	// some code goes here
	back := c.ll.Back()
	if back != nil {
		kv := back.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(kv.value.Len()) + int64(len(kv.key))
		c.ll.Remove(back)
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Len the number of cache entries
func (c *Cache) Len() int {
	// some code goes here
	return c.ll.Len()
}
