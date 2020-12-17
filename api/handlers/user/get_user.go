package user

import (
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

func NewGetUserHandler(r user.Repository) *GetUserHandler {
	return &GetUserHandler{
		userRepo: r,
	}
}

type GetUserHandler struct {
	userRepo user.Repository
}

// Handle ユーザーを取得します。GamersNoteIDで検索します。
func (h GetUserHandler) Handle(params o.GetUserParams) middleware.Responder {
	// u, err := h.usersRepo.GetUserByGamersNoteID(params.GamersnoteID)
	// if err != nil {
	// 	return o.NewGetUserDefault(500)
	// }
	// if u == nil {
	// 	return o.NewGetUserDefault(404)
	// }
	// return o.NewGetUserOK().WithPayload(u.ToModel())
	return o.NewGetUserOK()
}
