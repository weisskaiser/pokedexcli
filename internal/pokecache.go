package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	lock    sync.Mutex
	d       time.Duration
}

func NewCache(d time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		d:       d,
		lock:    sync.Mutex{},
	}
	c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.entries[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	go func() {
		t := time.NewTicker(c.d)
		for range t.C {
			c.lock.Lock()
			for k, v := range c.entries {
				if time.Since(v.createdAt) > c.d {
					delete(c.entries, k)
				}
			}
			c.lock.Unlock()
		}
	}()
}
