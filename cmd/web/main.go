// Package foobar does this.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	defaultPort = 8888
)

func run() error {
	app := newApplication()
	mux := chi.NewRouter()
	assetFileServer := http.FileServer(http.Dir("./assets"))

	mux.Use(middleware.Logger)

	mux.Handle("/css/*", assetFileServer)
	mux.Handle("/js/*", assetFileServer)

	mux.Get("/", app.root)

	// Paginating partials.
	mux.Get("/dashboard", app.dashboard)
	mux.Get("/news/{offset}/{limit}", app.newsPartial)
	mux.Get("/weather/{offset}/{limit}", app.weatherPartial)

	// Basic comboboxes.
	mux.Get("/comboboxes", app.comboboxes)
	mux.Get("/words/search/radio", app.wordsSearchRadio)
	mux.Get("/words/search/checkbox", app.wordsSearchCheckbox)

	// Echarts, and reloading.
	mux.Get("/echarts", app.echarts)

	mux.Get("/tabs", app.tabs)

	server := newServer(defaultPort, mux)

	log.Printf("Listening on :%d\n", defaultPort)

	err := server.ListenAndServe()

	return fmt.Errorf("server.ListenAndServe(): %w", err)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
