package like

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
	FindOne(like *dtos.Like) (bool, error)
	PutOne(like *dtos.Like) error
	DeleteOne(like *dtos.Like) error
}

// repository ユーザーのリポジトリ実装
type repository struct {
	db *gorm.DB
}
