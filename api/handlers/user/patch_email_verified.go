package user

import (
	"gamersnote.com/v1/repositories/user"
	o "gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils/tmpemail"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

func NewPatchEmailVefifiedHandler(r user.Repository, t tmpemail.Service) *PatchEmailVefifiedHandler {
	return &PatchEmailVefifiedHandler{
		userRepo:    r,
		tmpemailSvc: t,
	}
}

type PatchEmailVefifiedHandler struct {
	userRepo    user.Repository
	tmpemailSvc tmpemail.Service
}

// PatchEmailVefifiedHandler コードを検証し、メールアドレスを更新します
func (h PatchEmailVefifiedHandler) Handle(params o.PatchUserEmailVerifiedParams) middleware.Responder {
	uid := string(*params.Body.UID)

	// uidでキャッシュを検索
	data := h.tmpemailSvc.Get(uid)
	if data == nil {
		return o.NewPatchUserVerifiedDefault(400)
	}

	// codeを比較
	err := bcrypt.CompareHashAndPassword(data.HashedCode, []byte(*params.Body.Code))
	if err != nil {
		return o.NewPatchUserVerifiedDefault(400)
	}

	// DBに保存
	u, err := h.userRepo.GetOneByUserID(uid)
	if err != nil {
		return o.NewPutUserDefault(500)
	}
	u.Email = data.Email
	u, err = h.userRepo.UpdateOne(u)
	if err != nil {
		return o.NewPutUserDefault(500)
	}

	return o.NewPatchUserEmailVerifiedOK()
}
