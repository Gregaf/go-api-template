package mocks

import (
	"context"

	"github.com/gregaf/portfolio-backend/pkg/store"
)

type MockProjectStore struct {
	ListProjectsFn func(ctx context.Context) ([]store.Project, error)
}

// ProjectsList implements DbStore
func (mps *MockProjectStore) ProjectsList(ctx context.Context) ([]store.Project, error) {
	return mps.ListProjectsFn(ctx)
}

func (mps *MockProjectStore) ProjectGet(ctx context.Context)    {}
func (mps *MockProjectStore) ProjectCreate(ctx context.Context) {}
func (mps *MockProjectStore) ProjectUpdate(ctx context.Context) {}
func (mps *MockProjectStore) ProjectDelete(ctx context.Context) {}
