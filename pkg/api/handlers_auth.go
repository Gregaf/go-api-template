package api

import (
	"context"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

var oauthStateString = "pseudo-random"

func (as *APIServer) handleGoogleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := as.app.AuthConfig.AuthCodeURL(oauthStateString)

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func (as *APIServer) handleGoogleCallback() http.HandlerFunc {
	logrus.Warn("I have been called!")
	return func(w http.ResponseWriter, r *http.Request) {

		state := r.FormValue("state")
		code := r.FormValue("code")

		if state != oauthStateString {
			http.Error(w, "Invalid state", http.StatusBadRequest)
			return
		}

		token, err := as.app.AuthConfig.Exchange(context.TODO(), code)
		if err != nil {
			http.Error(w, "Failed to exchange token", http.StatusBadRequest)
			return
		}

		logrus.WithField("token", token).Debug("Token")

		res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			http.Error(w, "Failed to get user info", http.StatusBadRequest)
			return
		}

		defer res.Body.Close()
		contents, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusBadRequest)
			return
		}

		w.Write(contents)
	}
}
