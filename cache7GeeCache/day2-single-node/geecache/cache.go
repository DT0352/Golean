package geecache

import (
	"exacple/cache7GeeCache/day1lru/lru"
	"sync"
)

type cache struct {
	// 锁
	mu sync.Mutex
	// 缓存
	lru *lru.Cache
	//缓存大小
	cacheBytes int64
}

func (c *cache) add(key string, value Byteview) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}
func (c *cache) get(key string) (v Byteview, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(Byteview), ok
	}
	return
}
