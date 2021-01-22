package tmpemail

import "github.com/patrickmn/go-cache"

func (s service) Register(uid string, data Data) {
	s.c.Set(uid, data, cache.DefaultExpiration)
}
