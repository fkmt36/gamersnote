package user

import (
	"gamersnote.com/v1/models"
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils/session"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

func NewPatchSigninedHandler(r user.Repository, s session.Service) *PatchSigninedHandler {
	return &PatchSigninedHandler{
		userRepo:   r,
		sessionSvc: s,
	}
}

type PatchSigninedHandler struct {
	userRepo   user.Repository
	sessionSvc session.Service
}

// PatchSignined ログイン
func (h PatchSigninedHandler) Handle(params o.PatchUserSigninedParams) middleware.Responder {

	// DBからユーザーを取得
	u, err := h.userRepo.GetOneByEmail(string(params.Body.Email))
	if err != nil {
		return o.NewPatchUserSigninedDefault(500)
	}
	if u == nil {
		res := &models.Error{Message: "そのメールアドレスは登録されていません"}
		return o.NewPatchUserSigninedDefault(400).WithPayload(res)
	}

	// パスワードを検証
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(params.Body.Password))
	if err != nil {
		return o.NewPatchUserSigninedDefault(401)
	}

	// セッションをアタッチしてレスポンス
	res, err := attachSession(u.UserID, o.NewPatchUserVerifiedOK().WithPayload(u.ToModel()), h.sessionSvc)
	if err != nil {
		return o.NewPatchUserSigninedDefault(500)
	}
	return res
}
