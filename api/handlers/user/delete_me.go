package user

import (
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

func NewDeleteMeHandler(r user.Repository) *DeleteMeHandler {
	return &DeleteMeHandler{
		userRepo: r,
	}
}

type DeleteMeHandler struct {
	userRepo user.Repository
}

// Handle 自分を削除します。
func (h DeleteMeHandler) Handle(params o.DeleteUserParams) middleware.Responder {
	// err := h.usersRepo.DeleteUser(token.Uid)
	// if err != nil {
	// 	return o.NewGetUserDefault(500)
	// }
	return o.NewDeleteUserNoContent()
}
