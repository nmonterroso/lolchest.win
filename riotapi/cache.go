package riotapi

import (
	"time"

	gc "github.com/patrickmn/go-cache"
)

type Cache interface {
	GetOrSet(key string, duration time.Duration, onEmpty func() (interface{}, error)) (interface{}, error)
	Clear(key string)
}

type wrapperCache struct {
	cache *gc.Cache
}

func NewCache() Cache {
	return &wrapperCache{
		cache: gc.New(5*time.Minute, 5*time.Minute),
	}
}

func (w *wrapperCache) GetOrSet(key string, duration time.Duration, onEmpty func() (interface{}, error)) (interface{}, error) {
	if value, ok := w.cache.Get(key); ok {
		return value, nil
	}

	value, err := onEmpty()
	if err != nil {
		return nil, err
	}

	w.cache.Set(key, value, duration)
	return value, nil
}

func (w *wrapperCache) Clear(key string) {
	w.cache.Delete(key)
}
