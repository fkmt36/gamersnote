package like

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/repositories/like"
	o "gamersnote.com/v1/restapi/operations/like"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
)

func NewDeleteLikeHandler(r like.Repository, s ctxuid.Service) *DeleteLikeHandler {
	return &DeleteLikeHandler{
		likeRepo:  r,
		ctxuidSvc: s,
	}
}

type DeleteLikeHandler struct {
	likeRepo  like.Repository
	ctxuidSvc ctxuid.Service
}

// Handle いいねを削除します。
func (h DeleteLikeHandler) Handle(params o.DeleteLikeParams) middleware.Responder {
	// uidを取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewDeleteLikeDefault(401)
	}

	// いいねを削除
	like := &dtos.Like{ArticleID: params.ArticleID, UserID: uid}
	err := h.likeRepo.DeleteOne(like)
	if err != nil {
		return o.NewDeleteLikeDefault(500)
	}
	return o.NewDeleteLikeNoContent()
}
