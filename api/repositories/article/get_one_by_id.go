package article

import (
	"errors"

	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

// GetOneByID - ArticleIDで記事を検索します。
func (r repository) GetOneByID(articleID string) (*dtos.Article, error) {
	a := &dtos.Article{}
	q := &dtos.Article{ArticleID: articleID}
	result := r.db.Where(q).Joins("Author", "Comments").First(a)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return a, nil
}
