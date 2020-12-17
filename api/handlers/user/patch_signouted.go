package user

import (
	o "gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils/ctxuid"
	"gamersnote.com/v1/utils/session"
	"github.com/go-openapi/runtime/middleware"
)

func NewPatchSignoutedHandler(ctxuidService ctxuid.Service, sessionService session.Service) *PatchSignoutedHandler {
	return &PatchSignoutedHandler{
		ctxuidService:  ctxuidService,
		sessionService: sessionService,
	}
}

type PatchSignoutedHandler struct {
	ctxuidService  ctxuid.Service
	sessionService session.Service
}

// Handle ログアウト
func (h PatchSignoutedHandler) Handle(params o.PatchUserSignoutedParams) middleware.Responder {

	uid := h.ctxuidService.GetUID(params.HTTPRequest)
	if uid == "" {
		return o.NewPatchUserSignoutedDefault(401)
	}

	id, err := params.HTTPRequest.Cookie("sessionid")
	if err != nil {
		return o.NewPatchUserSignoutedDefault(500)
	}

	h.sessionService.Delete(id.Value)
	return o.NewPatchUserSignoutedOK()
}
