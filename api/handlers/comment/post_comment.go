package comment

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/repositories/comment"
	o "gamersnote.com/v1/restapi/operations/comment"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
)

// NewPostCommentHandler PostCommentHandlerのコンストラクタ
func NewPostCommentHandler(r comment.Repository, c ctxuid.Service) *PostCommentHandler {
	return &PostCommentHandler{
		commentRepo: r,
		ctxuidSvc:   c,
	}
}

// PostCommentHandler コメントの投稿
type PostCommentHandler struct {
	commentRepo comment.Repository
	ctxuidSvc   ctxuid.Service
}

// Handle コメントを投稿します
func (h PostCommentHandler) Handle(params o.PostCommentParams) middleware.Responder {
	// uidを取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewPostCommentDefault(401)
	}
	// コメントを追加
	commentID := uuid.New().String()
	c := dtos.Comment{
		CommentID:       commentID,
		ArticleID:       params.ArticleID,
		AuthorID:        uid,
		Body:            string(*params.Body.Body),
		ParentCommentID: commentID,
	}
	comment, err := h.commentRepo.AddOne(&c)
	if err != nil {
		return o.NewGetCommentsDefault(500)
	}
	return o.NewPostCommentCreated().WithPayload(comment.ToModel())
}
