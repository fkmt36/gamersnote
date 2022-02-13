package article

import (
	"sort"

	"gamersnote.com/v1/repositories/article"
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/article"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetArticleByUsernameHandler GetArticleByUsernameHandlerのコンストラクタ
func NewGetArticleByUsernameHandler(r article.Repository, u user.Repository) *GetArticleByUsernameHandler {
	return &GetArticleByUsernameHandler{
		articleRepo: r,
		userRepo:    u,
	}
}

// GetArticleByUsernameHandler usernameで記事を取得
type GetArticleByUsernameHandler struct {
	articleRepo article.Repository
	userRepo    user.Repository
}

// Handle Usernameで記事を検索して返します
func (h GetArticleByUsernameHandler) Handle(params o.GetTheUsersArticlesParams) middleware.Responder {
	// ユーザー取得
	u, err := h.userRepo.GetOneByUsername(params.Username)
	if err != nil {
		return o.NewGetTheUsersArticlesDefault(500)
	}
	if u == nil {
		return o.NewGetTheUsersArticlesDefault(404)
	}
	// 記事取得
	a, err := h.articleRepo.GetByUserID(u.UserID)
	if err != nil {
		return o.NewGetArticleDefault(500)
	}
	// 日付順にソート
	sort.Slice(a, func(i, j int) bool {
		return a[i].CreatedAt.After(a[j].CreatedAt)
	})
	return o.NewGetTheUsersArticlesOK().WithPayload(*a.ToModels())
}
