package httpserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"MTS/profile/httpserver/api"
)

func New() *http.Server {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	subApp := api.New()
	subApp.ApplyEndpoints(router)

	return &http.Server{
		Addr:    "localhost:3001",
		Handler: router,
	}
}
