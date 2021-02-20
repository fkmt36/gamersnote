package article

import "gamersnote.com/v1/dtos"

// DeleteOne 記事を削除します
func (r repository) DeleteOne(uid string, articleID string) error {
	q := &dtos.Article{AuthorID: uid, ArticleID: articleID}
	result := r.db.Delete(q)
	return result.Error
}
