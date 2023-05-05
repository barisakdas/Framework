package cache

import "time"

type ICacheService interface {
	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{}, expiration time.Duration)
	Remove(key interface{})
}
