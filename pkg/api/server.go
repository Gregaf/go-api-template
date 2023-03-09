package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gregaf/portfolio-backend/pkg/store"
	"github.com/sirupsen/logrus"
)

const DEFAULT_TIMEOUT = time.Second * 3

type APIServer struct {
	addr  string
	store *store.Store
}

func NewAPIServer(addr string, store *store.Store) (*APIServer, error) {
	if addr == "" {
		return nil, errors.New("address cannot be empty")
	}

	if store == nil {
		return nil, errors.New("store cannot be nil")
	}

	return &APIServer{addr: addr, store: store}, nil
}

func (s *APIServer) router() http.Handler {
	router := chi.NewRouter()

	// Setup middleware
	// Setup routes
	router.Use(middleware.Logger)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names, err := s.store.ListCollections(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var namesConcat string
		for _, name := range names {
			namesConcat += name + ", "
		}

		w.Write([]byte(namesConcat))
	})

	return router
}

func (s *APIServer) Start(stop <-chan struct{}) error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.router(),
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
