package repositories

import (
	"errors"

	"gamersnote.com/v1/configs"
	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

var usersRepository *UsersRepository

// GetUsersRepository usersRepositoryを返します。初回は初期化が行われます。
func GetUsersRepository() *UsersRepository {
	if usersRepository == nil {
		usersRepository = &UsersRepository{
			db: configs.GetDB(),
		}
	}
	return usersRepository
}

// UsersRepository ユーザーのリポジトリです。
type UsersRepository struct {
	db *gorm.DB
}

// GetUserByGamersNoteID GamersNoteIDによってユーザーを検索します。ユーザーが存在しない場合はnilを返します。
func (r *UsersRepository) GetUserByGamersNoteID(gamersNoteID string) (*dtos.User, error) {
	u := &dtos.User{}
	q := &dtos.User{GamersNoteID: gamersNoteID}
	result := r.db.Where(q).First(u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return u, nil
}

// AddUser ユーザーを追加します。
func (r *UsersRepository) AddUser(u *dtos.User) (*dtos.User, error) {
	result := r.db.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
