package article

import (
	"sort"

	"gamersnote.com/v1/repositories/article"
	o "gamersnote.com/v1/restapi/operations/article"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
)

func NewGetArticleByUserLikedHandler(r article.Repository, c ctxuid.Service) *GetArticleByUserLikedHandler {
	return &GetArticleByUserLikedHandler{
		articleRepo: r,
		ctxuidSvc:   c,
	}
}

type GetArticleByUserLikedHandler struct {
	articleRepo article.Repository
	ctxuidSvc   ctxuid.Service
}

// Handle Usernameで記事を検索して返します。
func (h GetArticleByUserLikedHandler) Handle(params o.GetLikedArticlesParams) middleware.Responder {
	// uidを取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewGetLikedArticlesDefault(401)
	}
	// 記事を取得
	a, err := h.articleRepo.GetByUserLiked(uid)
	if err != nil {
		return o.NewGetLikedArticlesDefault(500)
	}
	// 日付順にソート
	sort.Slice(a, func(i, j int) bool {
		return a[i].CreatedAt.After(a[j].CreatedAt)
	})
	return o.NewGetLikedArticlesOK().WithPayload(*a.ToModels())
}
