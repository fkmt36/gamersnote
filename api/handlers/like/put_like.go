package like

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/repositories/like"
	o "gamersnote.com/v1/restapi/operations/like"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
)

func NewPutLikeHandler(r like.Repository, s ctxuid.Service) *PutLikeHandler {
	return &PutLikeHandler{
		likeRepo:  r,
		ctxuidSvc: s,
	}
}

type PutLikeHandler struct {
	likeRepo  like.Repository
	ctxuidSvc ctxuid.Service
}

// Handle いいねを削除します。
func (h PutLikeHandler) Handle(params o.PutLikeParams) middleware.Responder {
	// uidを取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewPutLikeDefault(401)
	}

	// いいねを削除
	like := &dtos.Like{ArticleID: params.ArticleID, UserID: uid}
	err := h.likeRepo.PutOne(like)
	if err != nil {
		return o.NewPutLikeDefault(500)
	}
	return o.NewPutLikeOK()
}
