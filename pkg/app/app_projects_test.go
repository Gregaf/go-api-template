package app_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/gregaf/portfolio-backend/pkg/app"
	"github.com/gregaf/portfolio-backend/pkg/mocks"
	"github.com/gregaf/portfolio-backend/pkg/store"
)

type listProjectsFn func(ctx context.Context) ([]store.Project, error)

func listProjectsHelper(projects []store.Project, err error) listProjectsFn {
	return func(ctx context.Context) ([]store.Project, error) {
		return projects, err
	}
}

// Laying out intended testing structure
func TestProjectsList(t *testing.T) {

	// Prepare static lists of projects to provide to test cases

	emptyList := []store.Project{}
	populatedList := []store.Project{
		{
			ID:          "1",
			BlogID:      "1",
			Name:        "Project 1",
			Description: "Project 1 Description",
		},
	}

	tests := map[string]struct {
		listProjectsFn   listProjectsFn
		expectedProjects []store.Project
	}{
		"Empty project list": {
			listProjectsFn:   listProjectsHelper(emptyList, nil),
			expectedProjects: emptyList,
		},
		"Populated project list": {
			listProjectsFn:   listProjectsHelper(populatedList, nil),
			expectedProjects: populatedList,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			mockStore := store.DbStore{
				ProjectStore: &mocks.MockProjectStore{ListProjectsFn: tc.listProjectsFn},
				BlogStore:    nil,
				ProfileStore: nil,
			}

			app := app.NewApp(&mockStore)

			got, err := app.ProjectsList()

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tc.expectedProjects) {
				t.Fatalf("project lists not equal, expected %v, got %v", tc.expectedProjects, got)
			}
		})
	}
}
