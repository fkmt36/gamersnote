package comment

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/repositories/comment"
	o "gamersnote.com/v1/restapi/operations/comment"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

func NewDeleteCommentHandler(r comment.Repository, c ctxuid.Service) *DeleteCommentHandler {
	return &DeleteCommentHandler{
		commentRepo: r,
		ctxuidSvc:   c,
	}
}

type DeleteCommentHandler struct {
	commentRepo comment.Repository
	ctxuidSvc   ctxuid.Service
}

// Handle Usernameで記事を検索して返します。
func (h DeleteCommentHandler) Handle(params o.DeleteCommentParams) middleware.Responder {
	// uidを取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewDeleteCommentDefault(401)
	}
	// コメントを取得
	q := &dtos.Comment{
		CommentID: uuid.New().String(),
		AuthorID:  uid,
	}
	c, err := h.commentRepo.GetOne(q)
	if err != nil {
		return o.NewDeleteCommentDefault(500)
	}
	if c == nil {
		return o.NewDeleteCommentDefault(404)
	}
	// コメントを削除
	err = h.commentRepo.DeleteOne(params.CommentID)
	if err != nil {
		return o.NewDeleteCommentDefault(500)
	}
	return o.NewDeleteCommentNoContent()
}
