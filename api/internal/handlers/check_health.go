package handlers

import (
	"net/http"

	"github.com/benKapl/cvmaker-api/internal/respond"
)

func handlerCheckHealth(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}

	respond.WithJSON(w, http.StatusOK, response{
		Status: "ok",
	})
}
