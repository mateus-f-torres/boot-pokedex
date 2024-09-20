package cache

import (
	"sync"
	"time"
)

type Cache struct {
	store    map[string]cacheEntry
	mutex    *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	mux := &sync.Mutex{}

	c := Cache{
		store:    map[string]cacheEntry{},
		mutex:    mux,
		interval: interval,
	}

	go func() {
		t := time.Tick(interval)
		for {
			select {
			case <-t:
				c.reapLoop()
			}
		}
	}()

	return c
}

func (c Cache) Get(key string) ([]byte, bool) {
	var val []byte

	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.store[key]
	if ok {
		val = entry.val
	}

	return val, ok
}

func (c Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.store[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) reapLoop() {
	now := time.Now()

	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.store {
		elapsed := now.Sub(v.createdAt)
		if elapsed >= c.interval {
			delete(c.store, k)
		}
	}
}
