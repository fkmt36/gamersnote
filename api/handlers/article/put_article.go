package article

import (
	"gamersnote.com/v1/repositories/article"
	o "gamersnote.com/v1/restapi/operations/article"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
)

// NewPutArticleHandler PutArticleHandlerのコンストラクタ
func NewPutArticleHandler(r article.Repository, s ctxuid.Service) *PutArticleHandler {
	return &PutArticleHandler{
		articleRepo: r,
		ctxuidSvc:   s,
	}
}

// PutArticleHandler 記事の更新
type PutArticleHandler struct {
	articleRepo article.Repository
	ctxuidSvc   ctxuid.Service
}

// Handle 記事を更新します。
func (h PutArticleHandler) Handle(params o.PutArticleParams) middleware.Responder {
	// セッションからUID取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewPutArticleDefault(401)
	}
	// 記事をDBから取得
	a, err := h.articleRepo.GetOneByID(params.ArticleID)
	if err != nil {
		return o.NewPutArticleDefault(500)
	}
	// 記事を更新
	a.ThumbnailURL = string(params.Body.ThumbnailURL)
	a.Title = string(params.Body.Title)
	a.Body = string(params.Body.Body)
	a, err = h.articleRepo.UpdateOne(a)
	if err != nil {
		return o.NewPutArticleDefault(500)
	}

	return o.NewPutArticleOK().WithPayload(a.ToModel())
}
