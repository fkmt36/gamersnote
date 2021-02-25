package article

import (
	"fmt"

	"gamersnote.com/v1/dtos"
)

// GetLatest 最新順で記事を取得します。キーワードでLIKE検索します。
func (r repository) GetLatest(offset int, keyword string) (dtos.Articles, error) {
	articles := []*dtos.Article{}
	k := fmt.Sprintf("%%%s%%", keyword)
	results := r.db.Where("Title LIKE ? OR Body LIKE ?", k, k).Order("created_at desc").Limit(20).Offset(offset).Joins("Author", "Comments").Find(&articles)
	if results.Error != nil {
		return nil, results.Error
	}
	return articles, nil
}
