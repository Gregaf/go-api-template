package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (a *APIServer) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {

	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			logrus.WithError(err).Error("failed to encode response")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (a *APIServer) decode(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}
