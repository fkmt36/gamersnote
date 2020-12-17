package user

import (
	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Repository ユーザーリポジトリインタフェース
type Repository interface {
	AddOne(user *dtos.User) (*dtos.User, error)
	DeleteOne(userID string) error
	GetOneByEmail(email string) (*dtos.User, error)
	GetOneByUserID(userID string) (*dtos.User, error)
	GetOneByUsername(username string) (*dtos.User, error)
	UpdateOne(user *dtos.User) (*dtos.User, error)
}

// repository ユーザーのリポジトリ実装
type repository struct {
	db *gorm.DB
}
