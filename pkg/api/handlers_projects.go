package api

import (
	"net/http"
)

// handleProjectList returns a HTTP HandlerFunc that is responsible for
// returning a list of projects.
func (as *APIServer) handleProjectList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		projects, err := as.app.ProjectsList()
		if err != nil {
			as.respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		as.respond(w, r, projects, http.StatusOK)
	}
}

func (a *APIServer) handleProjectGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a *APIServer) handleProjectCreate() http.HandlerFunc {

	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Tags        string `json:"tags"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		projectPayload := &request{}

		err := a.decode(w, r, projectPayload)
		if err != nil {
			return
		}

		// process the request via business logic

	}
}

func (a *APIServer) handleProjectUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (a *APIServer) handleProjectDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
