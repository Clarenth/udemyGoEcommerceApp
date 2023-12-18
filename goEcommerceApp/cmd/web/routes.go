package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *appliction) routes() http.Handler {
	// mux is short for Multiplexer
	mux := chi.NewRouter()

	return mux
}
