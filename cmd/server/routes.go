package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *App) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(SetJSONHeader)

	prefix := app.Data.Prefix

	mux.Route("/"+prefix, func(r chi.Router) {
		r.Get("/", app.getAll)
		r.Get("/{id}", app.findByID)
	})

	return mux
}
