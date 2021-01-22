package user

import (
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils/email"
	"gamersnote.com/v1/utils/tmpsession"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sethvargo/go-password/password"
)

func NewPatchPasswordResetHandler(r user.Repository, t tmpsession.Service, e email.Sender) *PatchPasswordResetHandler {
	return &PatchPasswordResetHandler{
		userRepo:      r,
		tmpsessionSvc: t,
		emailSvc:      e,
	}
}

type PatchPasswordResetHandler struct {
	userRepo      user.Repository
	tmpsessionSvc tmpsession.Service
	emailSvc      email.Sender
}

// Handle パスワードリセット用のリンクを送信します
func (h PatchPasswordResetHandler) Handle(params o.PatchUserPasswordResetParams) middleware.Responder {
	email := string(params.Body.Email)

	// uidを取得
	user, err := h.userRepo.GetOneByEmail(email)
	if err != nil {
		o.NewPatchUserPasswordResetDefault(500)
	}
	uid := user.UserID

	// コード生成
	code, err := password.Generate(16, 8, 0, false, false)
	if err != nil {
		return o.NewPostUserDefault(500)
	}

	// キャッシュに保存
	h.tmpsessionSvc.Register(code, tmpsession.Data{
		UID: uid,
	})

	// メール送信
	err = h.emailSvc.SendPasswordResetEmail(
		email,
		code,
	)
	if err != nil {
		return o.NewPostUserDefault(500)
	}

	return o.NewPatchUserPasswordResetOK()
}
