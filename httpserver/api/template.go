package api

import (
	"MTS/profile/httpserver/serializer"
	"MTS/profile/templates"
	"io"
	"net/http"
)

func (a *App) uploadTemplate(w http.ResponseWriter, r *http.Request) {
	reader := r.Body

	postBody, err := io.ReadAll(reader)
	if err != nil {
		errMsg := serializer.SerializeResponseJSON(err)
		http.Error(w, string(errMsg), http.StatusInternalServerError)
		return
	}

	parameters, err := templates.ParseTemplate(postBody)
	if err != nil {
		errMsg := serializer.SerializeResponseJSON(err)
		http.Error(w, string(errMsg), http.StatusInternalServerError)
		return
	}

	msg := serializer.SerializeResponseJSON(parameters)
	w.Write(msg)
}
