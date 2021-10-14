package httpserver

import (
	"net/http"

	"github.com/go-chi/chi"

	"MTS/profile/httpserver/api"
	"MTS/profile/httpserver/middlewares"
)

func New() *http.Server {
	router := chi.NewRouter()
	router.Use(middlewares.Logging)

	subApp := api.New()
	subApp.ApplyEndpoints(router)

	return &http.Server{
		Addr:    "localhost:3001",
		Handler: router,
	}
}
