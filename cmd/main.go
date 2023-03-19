package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gregaf/portfolio-backend/pkg/api"
	"github.com/gregaf/portfolio-backend/pkg/app"
	"github.com/gregaf/portfolio-backend/pkg/logging"
	"github.com/gregaf/portfolio-backend/pkg/store"
	"github.com/pkg/errors"
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

	storeLogger := logging.NewLogger(logrus.DebugLevel, os.Stdout)

	// Example using MongoDB as a database
	_, err := store.NewStoreMongo(databaseURL, storeLogger)
	if err != nil {
		return errors.Wrap(err, "failed to initialize mongo store")
	}

	// Store supplied must implement interface
	store := &store.DbStore{
		ExampleOneStore:   nil,
		ExampleTwoStore:   nil,
		ExampleThreeStore: nil,
	}

	app := app.NewApp(store)

	apiLogger := logging.NewLogger(logrus.DebugLevel, os.Stdout)

	server, err := api.NewAPIServer(addr, app, apiLogger)
	if err != nil {
		panic(err)
	}

	server.Start(stopper)

	return nil
}
