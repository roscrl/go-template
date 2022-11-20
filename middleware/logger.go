package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// Logger is a middleware that logs the start and end of each request, along
// with some useful data about what was requested, what the response status was,
// and how long it took to return.
func Logger(l *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				l.Infow("request",
					"method", r.Method,
					"path", r.URL.Path,
					"status", ww.Status(),
					"timeTaken", time.Since(t1),
					"bytesWritten", ww.BytesWritten(),
					"protocol", r.Proto,
					"requestId", middleware.GetReqID(r.Context()))
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
