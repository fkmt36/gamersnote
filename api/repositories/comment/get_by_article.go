package comment

import (
	"errors"

	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

// GetByUserID - userIDで記事を検索します。
func (r repository) GetByArticle(articleID string) (dtos.Comments, error) {
	comments := []*dtos.Comment{}
	q := &dtos.Comment{ArticleID: articleID}
	result := r.db.Where(q).Joins("Author").Find(&comments)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return comments, nil
}
