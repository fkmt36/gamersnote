package comment

import (
	"errors"

	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

// GetByUserID - userIDで記事を検索します。
func (r repository) GetOne(comment *dtos.Comment) (*dtos.Comment, error) {
	c := &dtos.Comment{}
	result := r.db.Where(comment).Joins("Author").Find(&c)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return c, nil
}
