package article

import (
	"errors"

	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

// GetByUserID - userIDで記事を検索します。
func (r repository) GetByUserLiked(uid string) (dtos.Articles, error) {
	articles := []*dtos.Article{}
	q := &dtos.Like{UserID: uid}
	result := r.db.Where(q).Joins("Article").Find(&articles)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return articles, nil
}
