package handlers

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/repositories"
	"gamersnote.com/v1/restapi/operations/article"
	"gamersnote.com/v1/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

var articlesHandler *ArticlesHandler

// GetArticlesHandler articlesHandlerを返します。初回は初期化が行われます。
func GetArticlesHandler() *ArticlesHandler {
	if articlesHandler == nil {
		articlesHandler = &ArticlesHandler{
			articlesRepo: repositories.GetArticlesRepository(),
		}
	}
	return articlesHandler
}

// ArticlesHandler 記事のハンドラー
type ArticlesHandler struct {
	articlesRepo *repositories.ArticlesRepository
}

// GetAllArticles 全ての記事を返します。
func (h *ArticlesHandler) GetAllArticles(params article.GetArticlesParams) middleware.Responder {
	articles, err := h.articlesRepo.GetAllArticles()
	if err != nil {
		return article.NewGetArticlesDefault(500)
	}
	return article.NewGetArticlesOK().WithPayload(*articles.ToModels())
}

// PostArticle 記事を投稿します。
func (h *ArticlesHandler) PostArticle(params article.PostArticleParams, token *utils.TokenPayload) middleware.Responder {
	a := dtos.Article{
		ArticleID:    uuid.New().String(),
		AuthorID:     token.Uid,
		ThumbnailURL: *params.Body.ThumbnailURL,
		Title:        *params.Body.Title,
		Body:         *params.Body.Body,
		IsPublished:  *params.Body.IsPublished,
	}
	result, err := h.articlesRepo.AddArticle(a)
	if err != nil {
		return article.NewPostArticleDefault(500)
	}
	return article.NewPostArticleCreated().WithPayload(result.ToModel())
}
