package comment

import "gamersnote.com/v1/dtos"

// AddOne - 記事を追加します。追加した記事を返します。
func (r repository) AddOne(comment *dtos.Comment) (*dtos.Comment, error) {
	result := r.db.Create(comment)
	if result.Error != nil {
		return nil, result.Error
	}
	c, err := r.GetOne(comment)
	if err != nil {
		return nil, err
	}
	return c, nil
}
