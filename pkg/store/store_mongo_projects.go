package store

import "context"

func (ms *MongoStore) ProjectsList(ctx context.Context) ([]Project, error) {
	return []Project{}, nil
}

func (ms *MongoStore) ProjectGet(ctx context.Context) {}

func (ms *MongoStore) ProjectCreate(ctx context.Context) {}

func (ms *MongoStore) ProjectUpdate(ctx context.Context) {}

func (ms *MongoStore) ProjectDelete(ctx context.Context) {}
