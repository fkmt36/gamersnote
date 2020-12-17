package article

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/repositories/article"
	o "gamersnote.com/v1/restapi/operations/article"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

func NewPostArticleHandler(r article.Repository, s ctxuid.Service) *PostArticleHandler {
	return &PostArticleHandler{
		articleRepo: r,
		ctxuidSvc:   s,
	}
}

type PostArticleHandler struct {
	articleRepo article.Repository
	ctxuidSvc   ctxuid.Service
}

// Handle 記事を投稿します。
func (h PostArticleHandler) Handle(params o.PostArticleParams) middleware.Responder {
	// uidを取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewPostArticleDefault(401)
	}

	// DBに保存
	a := dtos.Article{
		ArticleID:    uuid.New().String(),
		AuthorID:     uid,
		ThumbnailURL: string(params.Body.ThumbnailURL),
		Title:        string(params.Body.Title),
		Body:         string(params.Body.Body),
	}
	result, err := h.articleRepo.AddOne(a)
	if err != nil {
		return o.NewPostArticleDefault(500)
	}
	return o.NewPostArticleCreated().WithPayload(result.ToModel())
}
