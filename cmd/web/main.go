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
	mux.Get("/dashboard", app.dashboard)
	mux.Get("/news/{offset}/{limit}", app.newsPartial)
	mux.Get("/weather/{offset}/{limit}", app.weatherPartial)

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
