package article

import (
	"sort"

	"gamersnote.com/v1/repositories/article"
	o "gamersnote.com/v1/restapi/operations/article"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetAllArticlesHandler コンストラクタ
func NewGetAllArticlesHandler(r article.Repository) *GetAllArticlesHandler {
	return &GetAllArticlesHandler{
		articleRepo: r,
	}
}

// GetAllArticlesHandler 全ての記事を返します。
type GetAllArticlesHandler struct {
	articleRepo article.Repository
}

// Handle 全ての記事を返します。
func (h GetAllArticlesHandler) Handle(params o.GetArticlesParams) middleware.Responder {
	articles, err := h.articleRepo.GetAll()
	if err != nil {
		return o.NewGetArticlesDefault(500)
	}
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].CreatedAt.After(articles[j].CreatedAt)
	})
	return o.NewGetArticlesOK().WithPayload(*articles.ToModels())
}
