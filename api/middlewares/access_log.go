package middlewares

import (
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

type captureResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewCaptureResponseWriter captureResponseWriterのコンストラクタ
func NewCaptureResponseWriter(w http.ResponseWriter) *captureResponseWriter {
	return &captureResponseWriter{w, http.StatusOK}
}

// WriteHeader 実装
func (lrw *captureResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// AccessLog アクセスログミドルウェア
func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Infof("[ACCESS] START %v %v%v\n", r.Method, r.Host, r.URL)

		lrw := NewCaptureResponseWriter(w)
		next.ServeHTTP(lrw, r)

		elapsed := time.Since(start)

		code := lrw.statusCode
		if code >= 500 {
			log.Errorf("[ACCESS] END %v %v %v %v\n", r.Method, code, r.URL, elapsed)
		} else if code >= 400 {
			log.Warnf("[ACCESS] END %v %v %v %v\n", r.Method, code, r.URL, elapsed)
		} else {
			log.Infof("[ACCESS] END %v %v %v %v\n", r.Method, code, r.URL, elapsed)
		}
	})
}
