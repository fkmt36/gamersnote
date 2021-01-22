package tmpemail

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func NewService(defaultExpiration time.Duration, cleanupInterval time.Duration) Service {
	return &service{
		c: cache.New(defaultExpiration, cleanupInterval),
	}
}

// Service 仮メールアドレスサービスインタフェース
type Service interface {
	Register(uid string, data Data)
	Get(uid string) *Data
}

type service struct {
	c *cache.Cache
}

// Data 仮メールアドレス用データ
type Data struct {
	HashedCode []byte
	Email      string
}
