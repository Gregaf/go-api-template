package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

const REQUEST_TIMEOUT = time.Second * 60

func (a *APIServer) logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		start := time.Now()
		defer func() {
			a.log.WithFields(logrus.Fields{
				"status":       ww.Status(),
				"method":       r.Method,
				"requestURI":   r.RequestURI,
				"remoteAddr":   r.RemoteAddr,
				"responseTime": time.Since(start).String(),
			}).Info("Incoming request")
		}()

		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}

func (a *APIServer) timeout(next http.Handler) http.Handler {
	return middleware.Timeout(REQUEST_TIMEOUT)(next)
}
