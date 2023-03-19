package store

import (
	"context"

	"github.com/gregaf/portfolio-backend/pkg/logging"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	db  *mongo.Database
	log *logging.Logger
}

func NewStoreMongo(databaseURL string, logger *logging.Logger) (*MongoStore, error) {
	if databaseURL == "" {
		return nil, errors.New("databaseURL cannot be empty")
	}

	mongoServerAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(databaseURL).SetServerAPIOptions(mongoServerAPI)
	opts.Auth.AuthSource = "admin"

	conn, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	return &MongoStore{db: conn.Database("example"), log: logger}, nil
}
