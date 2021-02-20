package article

import (
	"gamersnote.com/v1/repositories/article"
	o "gamersnote.com/v1/restapi/operations/article"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
)

func NewDeleteArticleHandler(r article.Repository, s ctxuid.Service) *DeleteArticleHandler {
	return &DeleteArticleHandler{
		articleRepo: r,
		ctxuidSvc:   s,
	}
}

type DeleteArticleHandler struct {
	articleRepo article.Repository
	ctxuidSvc   ctxuid.Service
}

// Handle 記事を投稿します。
func (h DeleteArticleHandler) Handle(params o.DeleteArticleParams) middleware.Responder {
	// uidを取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewDeleteArticleDefault(401)
	}

	// 記事を削除
	err := h.articleRepo.DeleteOne(uid, params.ArticleID)
	if err != nil {
		return o.NewDeleteArticleDefault(500)
	}
	return o.NewDeleteArticleNoContent()
}
