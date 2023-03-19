package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func (a *APIServer) setupRoutes() {
	a.log.Info("Setting up routes...")

	a.router.Use(a.logger)
	a.router.Use(a.timeout)

	a.router.Route("/api", func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {

			r.Get("/", a.handleVersion())

			r.Route("/auth", func(r chi.Router) {
				r.Route("/google", func(r chi.Router) {
					r.Get("/login", a.handleGoogleLogin())
					r.Get("/callback", a.handleGoogleCallback())
				})
			})

			r.Route("/example", func(r chi.Router) {

			})
		})
	})

	a.log.Info("Routes setup complete.")
	a.log.Info("Logging API routes")
	err := chi.Walk(a.router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		a.log.WithFields(logrus.Fields{
			"method": method,
			"route":  route,
		}).Info("Logging route")
		return nil
	})
	if err != nil {
		a.log.WithError(err).Error("Failed to log routes")
	}

}

func (a *APIServer) handleVersion() http.HandlerFunc {
	const version = "v1.0.0"
	return func(w http.ResponseWriter, r *http.Request) {
		a.respond(w, r, fmt.Sprintf("REST API, %s", version), http.StatusOK)
	}
}
