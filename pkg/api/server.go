// Package api provides the HTTP REST API for the portfolio backend housing
// a variety of HTTP specific logic.
package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/gregaf/portfolio-backend/pkg/app"
	"github.com/sirupsen/logrus"
)

const DEFAULT_TIMEOUT = time.Second * 3

// Bind to app layer
type APIServer struct {
	addr   string
	app    *app.App
	router *chi.Mux
}

func NewAPIServer(addr string, app *app.App) (*APIServer, error) {
	if addr == "" {
		return nil, errors.New("address cannot be empty")
	}

	// Must review whether router should be passed as dependency instead
	srv := &APIServer{addr: addr, app: app, router: chi.NewRouter()}

	srv.setupRoutes()

	chi.Walk(srv.router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		logrus.WithFields(logrus.Fields{
			"method":      method,
			"route":       route,
			"handler":     handler,
			"middlewares": middlewares,
		}).Info("Logging route...")
		return nil
	})

	return srv, nil
}

func (a *APIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func (s *APIServer) Start(stop <-chan struct{}) error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s,
	}

	go func() {
		logrus.WithField("addr", srv.Addr).Info("starting API server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), DEFAULT_TIMEOUT)
	defer cancel()

	logrus.WithField("timeout", DEFAULT_TIMEOUT).Info("shutting down API server")
	return srv.Shutdown(ctx)
}
