package cache

type CacheItem struct {
	Value      interface{}
	Expiration int64
}
