package article

import (
	"gamersnote.com/v1/dtos"
)

// UpdateOne - Articleを更新します
func (r repository) UpdateOne(article *dtos.Article) (*dtos.Article, error) {
	result := r.db.Save(article)
	if result.Error != nil {
		return nil, result.Error
	}
	return article, nil
}
