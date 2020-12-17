package dtos

import (
	"time"

	"gamersnote.com/v1/models"
	"github.com/go-openapi/strfmt"
)

// Comment CommentGORMモデル
type Comment struct {
	CommentID       string `gorm:"primaryKey"`
	ArticleID       string
	AuthorID        string
	Author          User `gorm:"foreignKey:AuthorID"`
	Body            string
	ParentCommentID string
	Replies         Comments `gorm:"foreignKey:ParentCommentID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Comments Commentの配列
type Comments []*Comment

// ToModel DBモデルをモデルに変換します。
func (c Comment) ToModel() *models.Comment {
	m := &models.Comment{
		Body:      &c.Body,
		CommentID: models.BaseID(c.CommentID),
		CreatedAt: strfmt.Date(c.CreatedAt),
		Replies:   *c.Replies.ToModels(),
	}
	m.Author.User = *c.Author.ToModel()
	return m
}

// ToModels DBモデルをモデルに変換します。
func (c Comments) ToModels() *[]*models.Comment {
	comments := make([]*models.Comment, len(c))
	for i, e := range c {
		comments[i] = e.ToModel()
	}
	return &comments
}
