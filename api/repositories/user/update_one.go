package user

import "gamersnote.com/v1/dtos"

// UpdateOne ユーザーを更新します。
func (r repository) UpdateOne(user *dtos.User) (*dtos.User, error) {
	result := r.db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
