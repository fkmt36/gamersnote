package article

import (
	"sort"

	"gamersnote.com/v1/repositories/article"
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/article"
	"github.com/go-openapi/runtime/middleware"
)

func NewGetArticleByUsernameHandler(r article.Repository, u user.Repository) *GetArticleByUsernameHandler {
	return &GetArticleByUsernameHandler{
		articleRepo: r,
		userRepo:    u,
	}
}

type GetArticleByUsernameHandler struct {
	articleRepo article.Repository
	userRepo    user.Repository
}

// Handle Usernameで記事を検索して返します。
func (h GetArticleByUsernameHandler) Handle(params o.GetTheUsersArticlesParams) middleware.Responder {
	u, err := h.userRepo.GetOneByUsername(params.Username)
	if err != nil {
		return o.NewGetTheUsersArticlesDefault(500)
	}
	if u == nil {
		return o.NewGetTheUsersArticlesDefault(404)
	}
	a, err := h.articleRepo.GetByUserID(u.UserID)
	if err != nil {
		return o.NewGetArticleDefault(500)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].CreatedAt.After(a[j].CreatedAt)
	})
	return o.NewGetTheUsersArticlesOK().WithPayload(*a.ToModels())
}
