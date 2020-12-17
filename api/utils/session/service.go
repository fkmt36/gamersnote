package session

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func NewService(defaultExpiration time.Duration, cleanupInterval time.Duration) Service {
	return &service{
		c: cache.New(defaultExpiration, cleanupInterval),
	}
}

// Service セッションサービスインタフェース
type Service interface {
	Create(key string, data Data)
	Get(key string) *Data
	Delete(key string)
}

type service struct {
	c *cache.Cache
}

// Data セッション用データ
type Data struct {
	HashedSession []byte
	UID           string
}
