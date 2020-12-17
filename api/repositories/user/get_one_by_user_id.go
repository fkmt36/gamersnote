package user

import (
	"errors"

	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

// GetOneByUserID UserIDによってユーザーを検索します。ユーザーが存在しない場合はnilを返します。
func (r repository) GetOneByUserID(userID string) (*dtos.User, error) {
	u := &dtos.User{}
	q := &dtos.User{UserID: userID}
	result := r.db.Where(q).First(u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return u, nil
}
