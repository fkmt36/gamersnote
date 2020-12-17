package article

import "gamersnote.com/v1/dtos"

// GetAll 全ての記事を返します。
func (r repository) GetAll() (dtos.Articles, error) {
	articles := []*dtos.Article{}
	results := r.db.Joins("Author", "Comments").Find(&articles)
	if results.Error != nil {
		return nil, results.Error
	}
	return articles, nil
}
