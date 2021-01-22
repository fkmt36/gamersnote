package tmpsession

import "github.com/patrickmn/go-cache"

func (s service) Register(code string, data Data) {
	s.c.Set(code, data, cache.DefaultExpiration)
}
