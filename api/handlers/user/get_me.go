package user

import (
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
)

func NewGetMeHandler(r user.Repository, s ctxuid.Service) *GetMeHandler {
	return &GetMeHandler{
		userRepo:  r,
		ctxuidSvc: s,
	}
}

type GetMeHandler struct {
	userRepo  user.Repository
	ctxuidSvc ctxuid.Service
}

// Handle 自分のユーザー情報を取得します。
func (h GetMeHandler) Handle(params o.GetMeParams) middleware.Responder {
	println("bbb")
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewGetUserDefault(401)
	}
	u, err := h.userRepo.GetOneByUserID(uid)
	if err != nil {
		return o.NewGetUserDefault(500)
	}
	if u == nil {
		return o.NewGetUserDefault(404)
	}
	return o.NewGetUserOK().WithPayload(u.ToModel())
}
