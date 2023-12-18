package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	// mux is short for Multiplexer
	mux := chi.NewRouter()

	mux.Get("/virtual-terminal", app.VirtualTerminal)

	return mux
}
