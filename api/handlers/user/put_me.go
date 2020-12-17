package user

import (
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

func NewPutMeHandler(r user.Repository) *PutMeHandler {
	return &PutMeHandler{
		userRepo: r,
	}
}

type PutMeHandler struct {
	userRepo user.Repository
}

// PutMe 自分のユーザー情報を更新します。
func (h PutMeHandler) PutMe(params o.PutUserParams) middleware.Responder {
	// u, err := h.usersRepo.GetUserByUserID(token.Uid)
	// if err != nil {
	// 	return user.NewPutUserDefault(500)
	// }
	// u.Username = *params.Body.Username
	// u.AvatarURL = *params.Body.AvatarURL
	// u.Message = *params.Body.Message
	// u, err = h.usersRepo.UpdateUser(u)
	// if err != nil {
	// 	return user.NewPutUserDefault(500)
	// }
	// return user.NewPutUserOK().WithPayload(u.ToModel())
	return o.NewPutUserOK()
}
