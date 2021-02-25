package article

import (
	"fmt"

	"gamersnote.com/v1/dtos"
)

// GetPopular いいねの数順で記事を取得します。キーワードでLIKE検索します。
func (r repository) GetPopular(offset int, keyword string) (dtos.Articles, error) {
	articles := []*dtos.Article{}
	k := fmt.Sprintf("%%%s%%", keyword)
	results := r.db.Where("Title LIKE ? OR Body LIKE ?", k, k).Order("like_count desc").Limit(20).Offset(offset).Joins("Author", "Comments").Find(&articles)
	if results.Error != nil {
		return nil, results.Error
	}
	return articles, nil
}
