package mocks

import (
	"context"
)

type MockProjectStore struct {
	TemplateOneFn func(ctx context.Context)
}

// TemplateOne implements DbStore
func (mps *MockProjectStore) TemplateOne(ctx context.Context) {
	mps.TemplateOneFn(ctx)
}
