package app_test

import (
	"context"
	"testing"

	"github.com/gregaf/portfolio-backend/pkg/app"
	"github.com/gregaf/portfolio-backend/pkg/mocks"
	"github.com/gregaf/portfolio-backend/pkg/store"
)

type templateOneFn func(ctx context.Context)

// This would accept the return type as parameters (High order function)
func templateOneHelper() templateOneFn {
	return func(ctx context.Context) {
		// return inputs as output to control test cases
		return
	}
}

// Laying out intended testing structure
func TestProjectsList(t *testing.T) {

	tests := map[string]struct {
		templateOneFn  templateOneFn
		expectedOutput string
	}{
		"Test 1": {
			templateOneFn:  templateOneHelper(),
			expectedOutput: "",
		},
		"Test 2": {
			templateOneFn:  templateOneHelper(),
			expectedOutput: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			mockStore := store.DbStore{
				ExampleOneStore: &mocks.MockProjectStore{TemplateOneFn: tc.templateOneFn},
			}

			app.NewApp(&mockStore)

			// Call relevant function

			// Assert output
		})
	}
}
