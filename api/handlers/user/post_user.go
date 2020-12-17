package user

import (
	"gamersnote.com/v1/models"
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils/email"
	"gamersnote.com/v1/utils/tmpuser"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sethvargo/go-password/password"

	"golang.org/x/crypto/bcrypt"
)

func NewPostUserHandler(r user.Repository, tmpuserService tmpuser.Service, emailService email.Sender) *PostUserHandler {
	return &PostUserHandler{
		userRepo:   r,
		tmpuserSvc: tmpuserService,
		emailSvc:   emailService,
	}
}

type PostUserHandler struct {
	userRepo   user.Repository
	tmpuserSvc tmpuser.Service
	emailSvc   email.Sender
}

// PostUser 新規のユーザーを追加します。
func (h PostUserHandler) Handle(params o.PostUserParams) middleware.Responder {

	// メールアドレスの重複チェック
	u, err := h.userRepo.GetOneByEmail(string(params.Body.Email))
	if err != nil {
		return o.NewPostUserDefault(500)
	}
	if u != nil {
		return o.NewPostUserConflict().WithPayload(
			&models.Error{Message: "すでに登録されているメールアドレスです。"})
	}

	// ユーザー名の重複チェック
	u, err = h.userRepo.GetOneByUsername(string(params.Body.Username))
	if err != nil {
		return o.NewPostUserDefault(500)
	}
	if u != nil {
		return o.NewPostUserConflict().WithPayload(
			&models.Error{Message: "すでに使われているユーザー名です。"})
	}

	// ユーザー名がキャッシュにあるか
	data := h.tmpuserSvc.Get(string(params.Body.Username))
	if data != nil {
		return o.NewPostUserConflict().WithPayload(
			&models.Error{Message: "すでに使われているユーザー名です。"})
	}

	// ランダムな文字列を生成し、ハッシュ化
	code, err := password.Generate(16, 8, 0, false, false)
	if err != nil {
		return o.NewPostUserDefault(500)
	}
	hashedCode, err := bcrypt.GenerateFromPassword([]byte(code), 8)
	if err != nil {
		return o.NewPostUserDefault(500)
	}

	// パスワードのハッシュ化
	hashedPass, err := bcrypt.GenerateFromPassword(
		[]byte(params.Body.Password), 8)
	if err != nil {
		return o.NewPostUserDefault(500)
	}

	// キャッシュに保存
	h.tmpuserSvc.Register(
		string(params.Body.Username),
		tmpuser.Data{
			HashedCode: hashedCode,
			Email:      string(params.Body.Email),
			HashedPass: hashedPass})
	// h.TempUserCache.AddOne(
	// 	string(params.Body.Username),
	// 	cache.TempUserData{
	// 		HashedCode: hashedCode,
	// 		Email:      string(params.Body.Email),
	// 		HashedPass: hashedPass})

	// - メール送信
	err = h.emailSvc.SendVerificationEmail(
		string(params.Body.Email),
		string(params.Body.Username),
		code)
	if err != nil {
		return o.NewPostUserDefault(500)
	}

	return o.NewPostUserCreated()
}
