package api

import (
	sendinglist "MTS/profile/sending_list"
	"net/http"

	"github.com/go-chi/chi"
)

// ListMap reflects username to sending list
type ListMap map[string]*sendinglist.SendingList

// App contains client for auth service and listmap
type App struct {
	client  http.Client
	listMap ListMap
}

// New returns the new App object
func New() *App {
	client := http.Client{
		Transport: nil,
		Jar:       nil,
		Timeout:   0,
	}

	lm := make(ListMap)

	return &App{
		client:  client,
		listMap: lm,
	}
}

// ApplyEndpoints setup router endpoints
func (a *App) ApplyEndpoints(r *chi.Mux) {
	r.Get("/i", a.identify)
	// endpoint for creating sendinglist for user
	r.Get("/receivers/{user_id}", a.createSendingList)
	// endpoint for uploading html template golang
	r.Post("/upload_template", a.uploadTemplate)
}
