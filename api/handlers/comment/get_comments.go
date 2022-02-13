package comment

import (
	"sort"

	"gamersnote.com/v1/repositories/comment"
	o "gamersnote.com/v1/restapi/operations/comment"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetCommentsHandler GetCommentsHandlerのコンストラクタ
func NewGetCommentsHandler(r comment.Repository) *GetCommentsHandler {
	return &GetCommentsHandler{
		commentRepo: r,
	}
}

// GetCommentsHandler コメントの取得
type GetCommentsHandler struct {
	commentRepo comment.Repository
}

// Handle 記事IDを指定してコメントを全て取得します
func (h GetCommentsHandler) Handle(params o.GetCommentsParams) middleware.Responder {
	// コメント取得
	c, err := h.commentRepo.GetByArticle(params.ArticleID)
	if err != nil {
		return o.NewGetCommentsDefault(500)
	}
	// 日付順にソート
	sort.Slice(c, func(i, j int) bool {
		return c[i].CreatedAt.After(c[j].CreatedAt)
	})
	return o.NewGetCommentsOK().WithPayload(*c.ToModels())
}
