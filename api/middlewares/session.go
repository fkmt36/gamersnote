package middlewares

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"gamersnote.com/v1/utils/ctxuid"
	"gamersnote.com/v1/utils/session"
)

// Session セッションミドルウェア
func Session(next http.Handler, sessionService session.Service, ctxuidService ctxuid.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// sessionidを取り出す
		c1, err := r.Cookie("sessionid")
		if c1 == nil || err != nil {
			next.ServeHTTP(w, r)
			return
		}

		// sessionを取り出す
		c2, err := r.Cookie("session")
		if c2 == nil || err != nil {
			next.ServeHTTP(w, r)
			return
		}

		// キャッシュからsessionidを検索
		data := sessionService.Get(c1.Value)
		if data == nil {
			next.ServeHTTP(w, r)
			return
		}

		// sessionを検証
		err = bcrypt.CompareHashAndPassword(data.HashedSession, []byte(c2.Value))
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		// contextにuidを入れる
		r = ctxuidService.SetUID(data.UID, r)
		next.ServeHTTP(w, r)
	})
}
