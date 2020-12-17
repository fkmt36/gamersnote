package user

import (
	"errors"

	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

// GetOneByUsername Usernameによってユーザーを検索します。ユーザーが存在しない場合はnilを返します。
func (r repository) GetOneByUsername(username string) (*dtos.User, error) {
	u := &dtos.User{}
	q := &dtos.User{Username: username}
	result := r.db.Where(q).First(u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return u, nil
}
