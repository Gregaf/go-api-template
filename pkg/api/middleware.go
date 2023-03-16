package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

const REQUEST_TIMEOUT = time.Second * 60

func (a *APIServer) logger(next http.Handler) http.Handler {
	return middleware.Logger(next)
}

func (a *APIServer) timeout(next http.Handler) http.Handler {
	return middleware.Timeout(REQUEST_TIMEOUT)(next)
}
