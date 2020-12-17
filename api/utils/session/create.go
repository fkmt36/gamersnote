package session

import "github.com/patrickmn/go-cache"

func (s service) Create(key string, data Data) {
	s.c.Set(key, data, cache.DefaultExpiration)
}
