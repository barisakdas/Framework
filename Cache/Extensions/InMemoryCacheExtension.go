package cache

import (
	. "github.com/barisakdas/Framework/Cache/Interfaces"
	. "github.com/barisakdas/Framework/Cache/Models"
	. "github.com/barisakdas/Framework/Cache/Services"
)

func NewInMemoryCache() ICacheService {
	_service := &InMemoryCacheService{
		Cache: make(map[interface{}]CacheItem),
	}
	go _service.CleanupExpiredItems()
	return _service
}
