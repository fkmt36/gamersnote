package user

import (
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils/ctxuid"
	"github.com/go-openapi/runtime/middleware"
)

func NewDeleteMeHandler(r user.Repository, c ctxuid.Service) *DeleteMeHandler {
	return &DeleteMeHandler{
		userRepo:  r,
		ctxuidSvc: c,
	}
}

type DeleteMeHandler struct {
	userRepo  user.Repository
	ctxuidSvc ctxuid.Service
}

// Handle 自分を削除します。
func (h DeleteMeHandler) Handle(params o.DeleteUserParams) middleware.Responder {
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewGetUserDefault(401)
	}

	err := h.userRepo.DeleteOne(uid)
	if err != nil {
		return o.NewGetUserDefault(500)
	}
	return o.NewDeleteUserNoContent()
}
