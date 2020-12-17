package dtos

import (
	"time"

	"gamersnote.com/v1/models"
	"github.com/go-openapi/strfmt"
)

// User UserGORMモデル
type User struct {
	UserID    string `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	Message   string
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Users Userポインタの配列
type Users []*User

// ToModel DBモデルをモデルに変換
func (u User) ToModel() *models.User {
	return &models.User{
		UserID:       models.BaseID(u.UserID),
		Username:     models.Username(u.Username),
		Email:        models.Email(u.Email),
		Message:      models.ProfileMessage(u.Message),
		AvatarURL:    models.ImgURL(u.AvatarURL),
		RegisteredAt: models.BaseDate(strfmt.Date(u.CreatedAt)),
	}
}

// ToModels DBモデルをモデルに変換
func (u Users) ToModels() *[]*models.User {
	users := make([]*models.User, len(u))
	for i, e := range u {
		users[i] = e.ToModel()
	}
	return &users
}
