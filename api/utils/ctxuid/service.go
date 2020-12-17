package ctxuid

import "net/http"

func NewService() Service {
	return &service{}
}

// Service contextからuidを取り出したり、セットしたりします。
type Service interface {
	GetUID(r *http.Request) string
	SetUID(uid string, r *http.Request) *http.Request
}

type service struct{}
