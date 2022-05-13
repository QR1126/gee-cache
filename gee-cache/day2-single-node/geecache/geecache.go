package geecache

import (
	"fmt"
	"log"
	"sync"
)

// A Group is a cache namespace and associated data loaded spread over
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

// A Getter loads data for a key.
type Getter interface {
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function.
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// NewGroup create a new instance of Group
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	// some code goes here
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	group := Group{
		name:   name,
		getter: getter,
		mainCache: cache{
			cacheBytes: cacheBytes,
		},
	}
	groups[name] = &group
	return &group
}

// GetGroup returns the named group previously created with NewGroup, or
// nil if there's no such group.
func GetGroup(name string) *Group {
	// some code goes here
	mu.RLock()
	defer mu.RUnlock()
	if group, ok := groups[name]; ok {
		return group
	}
	return nil
}

// Get value for a key from cache
func (g *Group) Get(key string) (ByteView, error) {
	// some code goes here
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	if val, ok := g.mainCache.get(key); ok {
		log.Println("[GeeCache] hit")
		return val, nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	// some code goes here
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	// some code goes here
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	val := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, val)
	return val, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	// some code goes here
	g.mainCache.add(key, value)
}
