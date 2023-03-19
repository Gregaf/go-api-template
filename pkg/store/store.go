package store

import "context"

type DbStore struct {
	ExampleOneStore
	ExampleTwoStore
	ExampleThreeStore
}

type ExampleOneStore interface {
	TemplateOne(ctx context.Context)
}

type ExampleTwoStore interface {
	TemplateTwo(ctx context.Context)
}

type ExampleThreeStore interface {
	TemplateThree(ctx context.Context)
}
