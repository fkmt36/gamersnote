package handlers

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/models"
	"gamersnote.com/v1/repositories"
	"gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
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
		UserID:       uuid.New().String(),
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
