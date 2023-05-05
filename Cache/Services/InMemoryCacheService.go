package cache

import (
	"sync"
	"time"

	. "github.com/barisakdas/Framework/Cache/Models"
)

type InMemoryCacheService struct {
	Cache map[interface{}]CacheItem
	Mutex sync.RWMutex
}

func (c *InMemoryCacheService) Get(key interface{}) (interface{}, bool) {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	item, found := c.Cache[key]
	if !found {
		return nil, false
	}

	if item.Expiration > 0 && time.Now().UnixNano() > item.Expiration {
		return nil, false
	}

	return item.Value, true
}

func (c *InMemoryCacheService) Set(key interface{}, value interface{}, expiration time.Duration) {
	var expirationTime int64

	if expiration > 0 {
		expirationTime = time.Now().Add(expiration).UnixNano()
	}

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Cache[key] = CacheItem{
		Value:      value,
		Expiration: expirationTime,
	}
}

func (c *InMemoryCacheService) Remove(key interface{}) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	delete(c.Cache, key)
}

func (c *InMemoryCacheService) CleanupExpiredItems() {
	for {
		<-time.After(time.Minute)

		c.Mutex.Lock()
		for key, item := range c.Cache {
			if item.Expiration > 0 && time.Now().UnixNano() > item.Expiration {
				delete(c.Cache, key)
			}
		}
		c.Mutex.Unlock()
	}
}
