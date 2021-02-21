package like

import "gamersnote.com/v1/dtos"

// DeleteOne ユーザーを追加します。
func (r repository) DeleteOne(like *dtos.Like) error {
	// いいね削除
	result := r.db.Delete(like)
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
	a.LikeCount = a.LikeCount - 1
	result = r.db.Save(a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
