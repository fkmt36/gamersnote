package like

import "gamersnote.com/v1/dtos"

// PutOne いいねを追加します。
func (r repository) PutOne(like *dtos.Like) error {
	// いいね追加
	result := r.db.Create(like)
	if result.Error != nil {
		return result.Error
	}
	// 記事取得
	a := &dtos.Article{}
	q := &dtos.Article{ArticleID: like.ArticleID}
	result = r.db.Where(q).First(a)
	if result.Error != nil {
		return result.Error
	}
	// 記事更新
	a.LikeCount = a.LikeCount + 1
	result = r.db.Save(a)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
