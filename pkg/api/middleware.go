package api

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

func (a *APIServer) logger(next http.Handler) http.Handler {
	return middleware.Logger(next)
}
