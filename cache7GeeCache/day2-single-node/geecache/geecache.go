package geecache

import (
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function.
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// A Group is a cache namespace and associated data loaded spread over
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mu     sync.Mutex
	groups = make(map[string]*Group)
)

// NewGroup create a new instance of Group
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{name: name, getter: getter, mainCache: cache{cacheBytes: cacheBytes}}
	groups[name] = g
	return g
}

// Get value for a key from cache
func (g *Group) Get(key string) (Byteview, error) {
	if key == "" {
		return Byteview{}, fmt.Errorf("key is required")
	}
	if v, ok := g.mainCache.get(key); ok {
		log.Panicln("[GeeCache] hit")
		return v, nil
	}
	return g.load(key)

}

func (g *Group) load(key string) (value Byteview, err error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (Byteview, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return Byteview{}, nil
	}
	value := Byteview{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value Byteview) {
	g.mainCache.add(key, value)
}

//GetGroup returns the named group previously created with NewGroup, or
// nil if there's no such group.
func GetGroup(name string) *Group {
	mu.Lock()
	defer mu.Unlock()
	g := groups[name]
	return g
}
