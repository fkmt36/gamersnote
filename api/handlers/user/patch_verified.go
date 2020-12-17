package user

import (
	"gamersnote.com/v1/dtos"
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils/session"
	"gamersnote.com/v1/utils/tmpuser"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func NewPatchVefifiedHandler(r user.Repository, tmpuserService tmpuser.Service, sessionService session.Service) *PatchVefifiedHandler {
	return &PatchVefifiedHandler{
		userRepo:   r,
		tmpuserSvc: tmpuserService,
		sessionSvc: sessionService,
	}
}

type PatchVefifiedHandler struct {
	userRepo   user.Repository
	tmpuserSvc tmpuser.Service
	sessionSvc session.Service
}

// PatchVefified コードを検証し、ユーザーを検証済みにします
func (h PatchVefifiedHandler) Handle(params o.PatchUserVerifiedParams) middleware.Responder {

	// usernameでキャッシュを検索
	data := h.tmpuserSvc.Get(string(*params.Body.Username))
	if data == nil {
		return o.NewPatchUserVerifiedDefault(400)
	}

	// codeを比較
	err := bcrypt.CompareHashAndPassword(data.HashedCode, []byte(*params.Body.Code))
	if err != nil {
		return o.NewPatchUserVerifiedDefault(400)
	}

	// DBに保存
	u, err := h.userRepo.AddOne(&dtos.User{
		UserID:    uuid.New().String(),
		Username:  string(*params.Body.Username),
		Email:     data.Email,
		Password:  string(data.HashedPass),
		Message:   "",
		AvatarURL: "",
	})
	if err != nil {
		return o.NewPatchUserVerifiedDefault(500)
	}

	// セッションをアタッチしてレスポンス
	res, err := attachSession(u.UserID, o.NewPatchUserVerifiedOK().WithPayload(u.ToModel()), h.sessionSvc)
	if err != nil {
		return o.NewPatchUserVerifiedDefault(500)
	}
	return res
}
