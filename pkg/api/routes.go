package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Projects:  projectID, blogID, name, description, image, tags, date created, date updated

// Optionally bind a blog to a project
// Blog: blogID, projectID, title, description, tags, date created, date updated

// Profile: profileID, name, email, image, files (binary data), links, date created, date updated

func (a *APIServer) setupRoutes() {
	a.router.Use(a.logger)

	a.router.Get("/", a.handleHelloWorld())

	a.router.Route("/api/v1", func(r chi.Router) {

		r.Route("/projects", func(r chi.Router) {
			// Create project
			// Delete project
			// Update project
			// Get project
			// List projects
		})

		// This must be structured to programatically build basic blogs
		r.Route("/blogs", func(r chi.Router) {
			// Create blog
			// Delete blog
			// Update blog
			// Get blog
			// List blogs
		})

		// This will only house my profile (This is for learning purposes rather than hardingcoding values)
		r.Route("/profile", func(r chi.Router) {
			// Create profile
			// Delete profile
			// Update profile
			// Get profile
			// List profiles
		})
	})

}

func (a *APIServer) handleHelloWorld() http.HandlerFunc {
	// Process some string...

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there world!"))
	}
}
