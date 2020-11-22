package dtos

import (
	"time"

	"gamersnote.com/v1/models"
	"github.com/go-openapi/strfmt"
)

// User UserGORMモデル
type User struct {
	Base
	UserID       string `gorm:"primaryKey"`
	GamersNoteID string `gorm:"unique"`
	Username     string
	Message      string
	AvatarURL    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Users Userポインタの配列
type Users []*User

// ToModel DBモデルをモデルに変換
func (u User) ToModel() *models.User {
	return &models.User{
		UserID:       u.UserID,
		GamersnoteID: &u.GamersNoteID,
		Username:     &u.Username,
		Message:      &u.Message,
		AvatarURL:    &u.AvatarURL,
		RegisteredAt: strfmt.Date(u.CreatedAt),
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
