package handlers

import (
	"net/http"

	"github.com/benKapl/cvmaker_api/internal/respond"
)

func (a *API) handlerReset(w http.ResponseWriter, r *http.Request) {
	if a.Platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset is only allowed in dev environment"))
		return
	}

	err := a.DB.DeleteUsers(r.Context())
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Couldn't reset database", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database reset"))
}
