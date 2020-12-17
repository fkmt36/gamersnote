package tmpuser

import "github.com/patrickmn/go-cache"

func (s service) Register(username string, data Data) {
	s.c.Set(username, data, cache.DefaultExpiration)
}
