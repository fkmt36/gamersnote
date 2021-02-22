package comment

import "gamersnote.com/v1/dtos"

// UpdateOne ユーザーを更新します。
func (r repository) UpdateOne(comment *dtos.Comment) (*dtos.Comment, error) {
	result := r.db.Save(comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}
