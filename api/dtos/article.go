package dtos

import (
	"time"

	"gamersnote.com/v1/models"
	"github.com/go-openapi/strfmt"
)

// Article ArticleGORMモデル
type Article struct {
	ArticleID    string `gorm:"primaryKey"`
	AuthorID     string
	Author       User `gorm:"foreignKey:AuthorID; constraint:OnUpdate:CASCADE"`
	ThumbnailURL string
	Title        string
	Body         string
	Comments     Comments `gorm:"foreignKey:ArticleID"`
	LikeCount    int64
	IsPublished  bool
	PublishedAt  time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Articles Articleポインタの配列
type Articles []*Article

// ToModel DBモデルをモデルに変換します。
func (a Article) ToModel() *models.Article {
	m := &models.Article{
		ArticleID:    models.BaseID(a.ArticleID),
		Body:         models.BaseHTML(a.Body),
		Comments:     *a.Comments.ToModels(),
		CreatedAt:    models.BaseDate(strfmt.Date(a.CreatedAt)),
		LikeCount:    models.BaseCount(a.LikeCount),
		ThumbnailURL: models.ImgURL(a.ThumbnailURL),
		Title:        models.ArticleTitle(a.Title),
		UpdatedAt:    models.BaseDate(strfmt.Date(a.UpdatedAt)),
	}
	m.Author.User = *a.Author.ToModel()
	return m
}

// ToModels DBモデルをモデルに変換します。
func (a Articles) ToModels() *[]*models.Article {
	articles := make([]*models.Article, len(a))
	for i, e := range a {
		articles[i] = e.ToModel()
	}
	return &articles
}
