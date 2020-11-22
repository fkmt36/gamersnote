package repositories

import (
	"gamersnote.com/v1/configs"
	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

var articlesRepository *ArticlesRepository

// GetArticlesRepository articlesRepositoryを返します。初回は初期化が行われます。
func GetArticlesRepository() *ArticlesRepository {
	if articlesRepository == nil {
		articlesRepository = &ArticlesRepository{
			db: configs.GetDB(),
		}
	}
	return articlesRepository
}

// ArticlesRepository 記事のリポジトリ
type ArticlesRepository struct {
	db *gorm.DB
}

// GetAllArticles 全ての記事を返します。
func (r *ArticlesRepository) GetAllArticles() (dtos.Articles, error) {
	articles := []*dtos.Article{}
	results := r.db.Joins("Author", "Comments").Find(&articles)
	if results.Error != nil {
		return nil, results.Error
	}
	return articles, nil
}

// AddArticle - 記事を追加します。追加した記事を返します。
func (r *ArticlesRepository) AddArticle(article dtos.Article) (*dtos.Article, error) {
	result := r.db.Create(&article)
	if result.Error != nil {
		return nil, result.Error
	}
	return &article, nil
}
