package handlers

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/models"
	"gamersnote.com/v1/repositories"
	"gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils"
	"github.com/go-openapi/runtime/middleware"
)

var usersHandler *UsersHandler

// GetUsersHandler usersHandlerを返します。初回は初期化が行われます。
func GetUsersHandler() *UsersHandler {
	if usersHandler == nil {
		usersHandler = &UsersHandler{
			usersRepo: repositories.GetUsersRepository(),
		}
	}
	return usersHandler
}

// UsersHandler ユーザーのハンドラー
type UsersHandler struct {
	usersRepo *repositories.UsersRepository
}

// PostUser 新規のユーザーを追加します。
func (h *UsersHandler) PostUser(params user.PostUserParams, token *utils.TokenPayload) middleware.Responder {
	u, err := h.usersRepo.GetUserByGamersNoteID(*params.Body.GamersnoteID)
	if err != nil {
		return user.NewPostUserDefault(500)
	}
	if u != nil {
		payload := &models.Error{Message: "すでに使われているGamersNoteIDです。"}
		return user.NewPostUserConflict().WithPayload(payload)
	}
	u = &dtos.User{
		UserID:       token.Uid,
		GamersNoteID: *params.Body.GamersnoteID,
		Message:      *params.Body.Message,
		Username:     *params.Body.Username,
	}
	result, err := h.usersRepo.AddUser(u)
	if err != nil {
		return user.NewPostUserDefault(500)
	}
	return user.NewPostUserCreated().WithPayload(result.ToModel())
}

// PutMe 自分のユーザー情報を更新します。
func (h *UsersHandler) PutMe(params user.PutUserParams, token *utils.TokenPayload) middleware.Responder {
	u, err := h.usersRepo.GetUserByUserID(token.Uid)
	if err != nil {
		return user.NewPutUserDefault(500)
	}
	u.GamersNoteID = *params.Body.GamersnoteID
	u.Username = *params.Body.Username
	u.AvatarURL = *params.Body.AvatarURL
	u.Message = *params.Body.Message
	u, err = h.usersRepo.UpdateUser(u)
	if err != nil {
		return user.NewPutUserDefault(500)
	}
	return user.NewPutUserOK().WithPayload(u.ToModel())
}

// GetUser ユーザーを取得します。GamersNoteIDで検索します。
func (h *UsersHandler) GetUser(params user.GetUserParams) middleware.Responder {
	u, err := h.usersRepo.GetUserByGamersNoteID(params.GamersnoteID)
	if err != nil {
		return user.NewGetUserDefault(500)
	}
	if u == nil {
		return user.NewGetUserDefault(404)
	}
	return user.NewGetUserOK().WithPayload(u.ToModel())
}

// GetMe 自分のユーザー情報を取得します。
func (h *UsersHandler) GetMe(params user.GetMeParams, token *utils.TokenPayload) middleware.Responder {
	u, err := h.usersRepo.GetUserByUserID(token.Uid)
	if err != nil {
		return user.NewGetUserDefault(500)
	}
	if u == nil {
		return user.NewGetUserDefault(404)
	}
	return user.NewGetUserOK().WithPayload(u.ToModel())
}

// DeleteMe 自分を削除します。
func (h *UsersHandler) DeleteMe(params user.DeleteUserParams, token *utils.TokenPayload) middleware.Responder {
	err := h.usersRepo.DeleteUser(token.Uid)
	if err != nil {
		return user.NewGetUserDefault(500)
	}
	return user.NewDeleteUserNoContent()
}
