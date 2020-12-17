package responder

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func NewSetCookieResponder(responder middleware.Responder) *SetCookieResponder {
	return &SetCookieResponder{
		responder: responder,
	}
}

type SetCookieResponder struct {
	responder middleware.Responder
	cookies   []http.Cookie
}

func (this *SetCookieResponder) WriteResponse(rw http.ResponseWriter, p runtime.Producer) {
	for _, c := range this.cookies {
		http.SetCookie(rw, &c)
	}
	this.responder.WriteResponse(rw, p)
}

func (this *SetCookieResponder) AddCookie(c http.Cookie) *SetCookieResponder {
	this.cookies = append(this.cookies, c)
	return this
}
