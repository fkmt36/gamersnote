package ctxuid

import (
	"context"
	"net/http"
)

// SetUID UIDをセット
func (s service) SetUID(uid string, r *http.Request) *http.Request {
	ctx := context.WithValue(r.Context(), Uidkey, uid)
	r = r.WithContext(ctx)
	return r
}
