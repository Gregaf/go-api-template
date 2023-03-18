package store

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	db *mongo.Database
}

func NewStoreMongo(databaseURL string) (*MongoStore, error) {

	mongoServerAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(databaseURL).SetServerAPIOptions(mongoServerAPI)
	opts.Auth.AuthSource = "admin"

	conn, err := mongo.Connect(context.TODO(), opts)
	if err != nil {

		return nil, errors.Wrap(err, "failed to connect to database")
	}

	return &MongoStore{db: conn.Database("portfolio")}, nil
}
