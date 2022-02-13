package article

import (
	"gamersnote.com/v1/repositories/article"
	o "gamersnote.com/v1/restapi/operations/article"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetArticleByArticleIDHandler GetArticleByArticleIDHandlerのコンストラクタ
func NewGetArticleByArticleIDHandler(r article.Repository) *GetArticleByArticleIDHandler {
	return &GetArticleByArticleIDHandler{
		articleRepo: r,
	}
}

// GetArticleByArticleIDHandler 記事IDから記事を取得
type GetArticleByArticleIDHandler struct {
	articleRepo article.Repository
}

// Handle ArticleIDで記事を検索して返します
func (h GetArticleByArticleIDHandler) Handle(params o.GetArticleParams) middleware.Responder {
	a, err := h.articleRepo.GetOneByID(params.ArticleID)
	if err != nil {
		return o.NewGetArticleDefault(500)
	}
	if a == nil {
		return o.NewGetArticleDefault(404)
	}
	return o.NewGetArticleOK().WithPayload(a.ToModel())
}
