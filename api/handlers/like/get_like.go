package like

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/repositories/like"
	o "gamersnote.com/v1/restapi/operations/like"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
)

func NewGetLikeHandler(r like.Repository, s ctxuid.Service) *GetLikeHandler {
	return &GetLikeHandler{
		likeRepo:  r,
		ctxuidSvc: s,
	}
}

type GetLikeHandler struct {
	likeRepo  like.Repository
	ctxuidSvc ctxuid.Service
}

// Handle いいねを削除します。
func (h GetLikeHandler) Handle(params o.GetLikeParams) middleware.Responder {
	// uidを取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewGetLikeUnauthorized()
	}

	// いいねを取得
	like := &dtos.Like{ArticleID: params.ArticleID, UserID: uid}
	found, err := h.likeRepo.FindOne(like)
	if err != nil {
		return o.NewGetLikeDefault(500)
	}
	if !found {
		return o.NewGetLikeNotFound()
	}
	return o.NewGetLikeOK()
}
