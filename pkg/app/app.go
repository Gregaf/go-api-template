package app

import (
	"os"

	"github.com/gregaf/portfolio-backend/pkg/store"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type App struct {
	AuthConfig *oauth2.Config
	store      *store.DbStore
}

func NewApp(store *store.DbStore) *App {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/v1/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	return &App{store: store, AuthConfig: googleOauthConfig}
}
