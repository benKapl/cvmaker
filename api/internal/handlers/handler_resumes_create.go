package handlers

import "net/http"

func (a *API) handlerResumesCreate(w http.ResponseWriter, r *http.Request) {
	a.ResumeService.CreateResume()

}
