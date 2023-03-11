package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gregaf/portfolio-backend/pkg/api"
	"github.com/gregaf/portfolio-backend/pkg/store"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	stopper := make(chan struct{})
	go func() {
		<-done
		close(stopper)
	}()

	addr := os.Getenv("API_ADDR")
	databaseURL := os.Getenv("DATABASE_URL")

	store, err := store.NewStore(databaseURL)
	if err != nil {
		panic(err)
	}

	server, err := api.NewAPIServer(addr, store)
	if err != nil {
		panic(err)
	}

	server.Start(stopper)

	return nil
}
