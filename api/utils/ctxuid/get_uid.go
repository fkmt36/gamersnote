package ctxuid

import "net/http"

// GetUID UIDを取得
func (s service) GetUID(r *http.Request) string {
	uid, ok := r.Context().Value(Uidkey).(string)
	if !ok {
		return ""
	}
	return uid
}
