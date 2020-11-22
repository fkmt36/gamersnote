package utils

import (
	"github.com/go-openapi/errors"
)

type TokenPayload struct {
	Uid string
}

func AuthHandler(token string) (*TokenPayload, error) {
	if token != "" {
		println(token)
		payload := &TokenPayload{
			Uid: token,
		}
		return payload, nil
	}
	return nil, errors.Unauthenticated("invalid token")
}
