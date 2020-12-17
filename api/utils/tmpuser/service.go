package tmpuser

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func NewService(defaultExpiration time.Duration, cleanupInterval time.Duration) Service {
	return &service{
		c: cache.New(defaultExpiration, cleanupInterval),
	}
}

// Service 仮登録ユーザーサービスインタフェース
type Service interface {
	Register(username string, data Data)
	Get(username string) *Data
}

type service struct {
	c *cache.Cache
}

// Data 仮登録ユーザー用データ
type Data struct {
	HashedCode []byte
	Email      string
	HashedPass []byte
}
