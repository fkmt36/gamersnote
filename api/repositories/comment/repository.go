package comment

import (
	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Repository 記事のリポジトリインタフェース
type Repository interface {
	AddOne(comment *dtos.Comment) (*dtos.Comment, error)
	GetOne(comment *dtos.Comment) (*dtos.Comment, error)
	GetByArticle(articleID string) (dtos.Comments, error)
	UpdateOne(comment *dtos.Comment) (*dtos.Comment, error)
	DeleteOne(commentID string) error
}

// ArticleRepository 記事のリポジトリ実装
type repository struct {
	db *gorm.DB
}
