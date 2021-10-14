package api

import (
	"MTS/profile/httpserver/serializer"
	"errors"
	"io"
	"net/http"
)

func (a *App) identify(w http.ResponseWriter, r *http.Request) {
	// check for access cookie
	_, err := r.Cookie("refresh")

	if err != nil {
		err := errors.New("You are not logged in!")
		http.Error(w, string(serializer.SerializeResponseJSON(err)), http.StatusForbidden)
		return
	}
	// generating request for a.Client
	newRequest, err := http.NewRequest("GET", "http://localhost:3000/me", nil)
	if err != nil {
		err := serializer.SerializeResponseJSON(err)
		http.Error(w, string(err), http.StatusBadRequest)
		return
	}

	// setting cookies to the request
	cookies := r.Cookies()
	for _, cookie := range cookies {
		newRequest.AddCookie(cookie)
	}

	// making request
	response, err := a.client.Do(newRequest)
	if err != nil {
		err := serializer.SerializeResponseJSON(err)
		http.Error(w, string(err), http.StatusBadGateway)
		return
	}

	if len(response.Cookies()) == 0 {
		err := errors.New("You are not logged in!")
		http.Error(w, string(serializer.SerializeResponseJSON(err)), http.StatusForbidden)
		return
	}

	// updating cookies
	for _, cookie := range response.Cookies() {
		http.SetCookie(w, cookie)
	}

	// reading data from the response
	data, err := io.ReadAll(response.Body)
	if err != nil {
		err := serializer.SerializeResponseJSON(err)
		http.Error(w, string(err), http.StatusInternalServerError)
		return
	}

	w.Write(serializer.SerializeResponseJSON(string(data)))
}
