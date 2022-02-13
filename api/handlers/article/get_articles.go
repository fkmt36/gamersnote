package article

import (
	"gamersnote.com/v1/repositories/article"
	o "gamersnote.com/v1/restapi/operations/article"
	"gamersnote.com/v1/utils"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetArticlesHandler コンストラクタ
func NewGetArticlesHandler(r article.Repository) *GetArticlesHandler {
	return &GetArticlesHandler{
		articleRepo: r,
	}
}

// GetArticlesHandler 記事を取得して返します。
type GetArticlesHandler struct {
	articleRepo article.Repository
}

// Handle 記事を取得します。offset, keyword, sortを指定してください。
func (h GetArticlesHandler) Handle(params o.GetArticlesParams) middleware.Responder {
	// ソート順 popular: 人気、その他: 最新
	var order string
	if params.Order != nil {
		order = *params.Order
	} else {
		order = ""
	}
	// 何記事目から取得するか offset+1記事目から取得します(0スタート)。
	var offset int
	if params.Offset != nil {
		offset = utils.Int64ToInt(*params.Offset)
	} else {
		offset = -1
	}
	// 検索する文字列
	var keyword string
	if params.Keyword != nil {
		keyword = *params.Keyword
	} else {
		keyword = ""
	}

	if order == "popular" {
		articles, err := h.articleRepo.GetPopular(offset, keyword)
		if err != nil {
			return o.NewGetArticlesDefault(500)
		}
		return o.NewGetArticlesOK().WithPayload(*articles.ToModels())
	}

	articles, err := h.articleRepo.GetLatest(offset, keyword)
	if err != nil {
		return o.NewGetArticlesDefault(500)
	}
	return o.NewGetArticlesOK().WithPayload(*articles.ToModels())
}
