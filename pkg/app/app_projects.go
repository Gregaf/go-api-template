package app

import (
	"context"

	"github.com/gregaf/portfolio-backend/pkg/store"
)

func (a *App) ProjectsList() ([]store.Project, error) {
	return a.store.ProjectStore.ProjectsList(context.TODO())
}
