package user

import (
	"errors"

	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

// GetOneByEmail Emailによってユーザーを検索します。ユーザーが存在しない場合はnilを返します。
func (r repository) GetOneByEmail(email string) (*dtos.User, error) {
	u := &dtos.User{}
	q := &dtos.User{Email: email}
	result := r.db.Where(q).First(u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return u, nil
}
