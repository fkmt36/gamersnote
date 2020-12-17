package user

import (
	"net/http"
	"time"

	"gamersnote.com/v1/utils/responder"
	"gamersnote.com/v1/utils/session"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func attachSession(userID string, res middleware.Responder, sessionService session.Service) (middleware.Responder, error) {
	// ランダムな文字列を生成
	s, err := password.Generate(16, 8, 0, false, false)
	if err != nil {
		return nil, err
	}

	// ハッシュ化
	hashedSession, err := bcrypt.GenerateFromPassword([]byte(s), 8)
	if err != nil {
		return nil, err
	}

	// uuidを生成
	uuid := uuid.New().String()

	// セッションストアに保存
	sessionService.Create(uuid, session.Data{HashedSession: hashedSession, UID: userID})
	// h.SessionCache.AddOne(
	// 	uuid, cache.SessionData{HashedSession: hashedSession, UserID: userID})

	// レスポンス
	c1 := http.Cookie{
		Name:     "sessionid",
		Value:    uuid,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().AddDate(10, 0, 0),
	}
	c2 := http.Cookie{
		Name:     "session",
		Value:    s,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().AddDate(10, 0, 0),
	}
	return responder.NewSetCookieResponder(res).AddCookie(c1).AddCookie(c2), nil
}
