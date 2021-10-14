package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

type App struct {
	client http.Client
}

// New returns the new App object
func New() *App {
	client := http.Client{
		Transport: nil,
		Jar:       nil,
		Timeout:   0,
	}

	return &App{
		client: client,
	}
}

func (a *App) ApplyEndpoints(r *chi.Mux) {
	r.Get("/i", a.identify)
}
