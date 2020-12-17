package user

import "gamersnote.com/v1/dtos"

// AddOne ユーザーを追加します。
func (r repository) AddOne(user *dtos.User) (*dtos.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
