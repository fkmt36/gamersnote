package utils

import (
	"context"

	"gamersnote.com/v1/configs"
	"github.com/go-openapi/errors"
)

type TokenPayload struct {
	Uid string
}

func AuthHandler(token string) (*TokenPayload, error) {
	ctx := context.Background()
	client, err := configs.GetFirebaseApp().Auth(ctx)
	if err != nil {
		return nil, errors.Unauthenticated("invalid token")
	}
	result, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, errors.Unauthenticated("invalid token")
	}
	if result == nil {
		return nil, errors.Unauthenticated("invalid token")
	}
	payload := &TokenPayload{
		Uid: result.UID,
	}
	return payload, nil
}
