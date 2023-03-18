package store

import (
	"context"
)

type DbStore struct {
	ProjectStore
	BlogStore
	ProfileStore
}

type Project struct {
	ID          string
	BlogID      string
	Name        string
	Description string
	Tags        []string
	CreatedDate int64
	UpdatedDate int64
}

type ProjectStore interface {
	ProjectsList(ctx context.Context) ([]Project, error)
	ProjectGet(ctx context.Context)
	ProjectCreate(ctx context.Context)
	ProjectUpdate(ctx context.Context)
	ProjectDelete(ctx context.Context)
}

type BlogStore interface {
	ListBlogs(ctx context.Context)
	GetBlog(ctx context.Context)
	CreateBlog(ctx context.Context)
	UpdateBlog(ctx context.Context)
	DeleteBlog(ctx context.Context)
}

type ProfileStore interface {
	GetProfile(ctx context.Context)
	CreateProfile(ctx context.Context)
	UpdateProfile(ctx context.Context)
}
