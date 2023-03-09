package store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	db *mongo.Database
}

func (s *Store) ListCollections(ctx context.Context) ([]string, error) {
	collections, err := s.db.ListCollectionNames(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list collections")
	}

	return collections, nil
}

func NewStore(databaseURL string) (*Store, error) {

	mongoServerAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(databaseURL).SetServerAPIOptions(mongoServerAPI)
	opts.Auth.AuthSource = "admin"

	conn, err := mongo.Connect(context.TODO(), opts)
	if err != nil {

		return nil, errors.Wrap(err, "failed to connect to database")
	}

	res, err := conn.Database("portfolio").Collection("skills").InsertOne(context.TODO(), map[string]interface{}{"name": "Go", "level": 5})
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert skill")
	}

	logrus.WithField("id", res.InsertedID).Info("Inserted skill")

	conn.Database("portfolio").CreateCollection(context.TODO(), "skills")

	return &Store{db: conn.Database("portfolio")}, nil
}
