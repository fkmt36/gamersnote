package article

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
	AddOne(article dtos.Article) (*dtos.Article, error)
	GetAll() (dtos.Articles, error)
	GetOneByID(articleID string) (*dtos.Article, error)
	GetByUserID(username string) (dtos.Articles, error)
	UpdateOne(article *dtos.Article) (*dtos.Article, error)
	DeleteOne(uid string, articleID string) error
}

// ArticleRepository 記事のリポジトリ実装
type repository struct {
	db *gorm.DB
}
