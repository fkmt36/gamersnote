package tmpsession

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func NewService(defaultExpiration time.Duration, cleanupInterval time.Duration) Service {
	return &service{
		c: cache.New(defaultExpiration, cleanupInterval),
	}
}

// Service 仮セッションサービスインタフェース
type Service interface {
	Register(code string, data Data)
	Get(code string) *Data
}

type service struct {
	c *cache.Cache
}

// Data 仮セッション用データ
type Data struct {
	UID string
}
