package article

import "gamersnote.com/v1/dtos"

// AddOne - 記事を追加します。追加した記事を返します。
func (r repository) AddOne(article dtos.Article) (*dtos.Article, error) {
	result := r.db.Create(&article)
	if result.Error != nil {
		return nil, result.Error
	}
	return &article, nil
}
