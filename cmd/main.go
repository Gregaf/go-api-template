package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gregaf/portfolio-backend/pkg/api"
	"github.com/gregaf/portfolio-backend/pkg/app"
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

	storeMongo, err := store.NewStoreMongo(databaseURL)
	if err != nil {
		panic(err)
	}

	store := &store.DbStore{
		ProjectStore: storeMongo,
		BlogStore:    nil,
		ProfileStore: nil,
	}

	app := app.NewApp(store)

	server, err := api.NewAPIServer(addr, app)
	if err != nil {
		panic(err)
	}

	server.Start(stopper)

	return nil
}
