package lru

import "container/list"

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
	return nil
}

// Add adds a value to the cache.
func (c *Cache) Add(key string, value Value) {
	// some code goes here
}

// Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	// some code goes here
	return
}

// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	// some code goes here
}

// Len the number of cache entries
func (c *Cache) Len() int {
	// some code goes here
	return 0
}
