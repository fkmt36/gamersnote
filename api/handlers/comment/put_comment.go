package comment

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/repositories/comment"
	o "gamersnote.com/v1/restapi/operations/comment"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

func NewPutCommentHandler(r comment.Repository, c ctxuid.Service) *PutCommentHandler {
	return &PutCommentHandler{
		commentRepo: r,
		ctxuidSvc:   c,
	}
}

type PutCommentHandler struct {
	commentRepo comment.Repository
	ctxuidSvc   ctxuid.Service
}

// Handle Usernameで記事を検索して返します。
func (h PutCommentHandler) Handle(params o.PutCommentParams) middleware.Responder {
	// uidを取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewPutCommentDefault(401)
	}
	// コメントを取得
	q := &dtos.Comment{
		CommentID: uuid.New().String(),
		AuthorID:  uid,
	}
	c, err := h.commentRepo.GetOne(q)
	if err != nil {
		return o.NewPutCommentDefault(500)
	}
	if c == nil {
		return o.NewPutCommentDefault(404)
	}
	// コメントを更新
	c.Body = string(*params.Body.Body)
	_, err = h.commentRepo.UpdateOne(c)
	if err != nil {
		return o.NewPutCommentDefault(500)
	}
	return o.NewPutCommentOK()
}
