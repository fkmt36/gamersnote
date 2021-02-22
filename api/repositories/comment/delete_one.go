package comment

import "gamersnote.com/v1/dtos"

// DeleteOne 記事を削除します
func (r repository) DeleteOne(commentID string) error {
	q := &dtos.Comment{CommentID: commentID}
	result := r.db.Delete(q)
	return result.Error
}
