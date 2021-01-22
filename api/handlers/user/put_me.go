package user

import (
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils/ctxuid"
	"gamersnote.com/v1/utils/email"
	"gamersnote.com/v1/utils/tmpemail"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func NewPutMeHandler(r user.Repository, s ctxuid.Service, t tmpemail.Service, e email.Sender) *PutMeHandler {
	return &PutMeHandler{
		userRepo:    r,
		ctxuidSvc:   s,
		tmpemailSvc: t,
		emailSvc:    e,
	}
}

type PutMeHandler struct {
	userRepo    user.Repository
	ctxuidSvc   ctxuid.Service
	tmpemailSvc tmpemail.Service
	emailSvc    email.Sender
}

// PutMe 自分のユーザー情報を更新します。
func (h PutMeHandler) Handle(params o.PutUserParams) middleware.Responder {
	// セッションからUID取得
	uid := h.ctxuidSvc.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewPutUserDefault(401)
	}
	// ユーザーをDBから取得
	u, err := h.userRepo.GetOneByUserID(uid)
	if err != nil {
		return o.NewPutUserDefault(500)
	}
	// Email以外アップデート
	u.Username = string(params.Body.Username)
	u.AvatarURL = string(params.Body.AvatarURL)
	u.Message = string(params.Body.Message)
	u, err = h.userRepo.UpdateOne(u)
	if err != nil {
		return o.NewPutUserDefault(500)
	}
	// Emailの処理
	if u.Email != string(params.Body.Email) {
		// ランダムな文字列を生成し、ハッシュ化
		code, err := password.Generate(16, 8, 0, false, false)
		if err != nil {
			return o.NewPostUserDefault(500)
		}
		hashedCode, err := bcrypt.GenerateFromPassword([]byte(code), 8)
		if err != nil {
			return o.NewPostUserDefault(500)
		}
		// キャッシュに保存
		h.tmpemailSvc.Register(uid, tmpemail.Data{
			HashedCode: hashedCode,
			Email:      string(params.Body.Email),
		})
		// メール送信
		err = h.emailSvc.SendEmailVerificationEmail(
			string(params.Body.Email),
			string(u.Username),
			uid,
			code)
		if err != nil {
			return o.NewPostUserDefault(500)
		}
	}

	return o.NewPutUserOK().WithPayload(u.ToModel())
}
