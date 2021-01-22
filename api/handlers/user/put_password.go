package user

import (
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils/ctxuid"
	"gamersnote.com/v1/utils/tmpsession"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

func NewPutPasswordHandler(r user.Repository, s ctxuid.Service, t tmpsession.Service) *PutPasswordHandler {
	return &PutPasswordHandler{
		userRepo:      r,
		ctxuidSvc:     s,
		tmpsessionSvc: t,
	}
}

type PutPasswordHandler struct {
	userRepo      user.Repository
	ctxuidSvc     ctxuid.Service
	tmpsessionSvc tmpsession.Service
}

// PutMe 自分のユーザー情報を更新します。
func (h PutPasswordHandler) Handle(params o.PutPasswordParams) middleware.Responder {
	p := string(params.Body.Password)
	c := params.Body.Code

	// UID取得
	var uid string = ""
	if c != "" {
		data := h.tmpsessionSvc.Get(c)
		if data == nil {
			return o.NewPutPasswordDefault(400)
		}
		uid = data.UID
	} else {
		uid = h.ctxuidSvc.GetUID(params.HTTPRequest)
		if uid == "" {
			return o.NewPutPasswordDefault(401)
		}
	}

	// ユーザー取得
	u, err := h.userRepo.GetOneByUserID(uid)
	if err != nil {
		return o.NewPutPasswordDefault(500)
	}

	// パスワード更新
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(p), 8)
	if err != nil {
		return o.NewPutPasswordDefault(500)
	}
	u.Password = string(hashedPass)
	u, err = h.userRepo.UpdateOne(u)
	if err != nil {
		return o.NewPutPasswordDefault(500)
	}

	return o.NewPutPasswordOK()
}
