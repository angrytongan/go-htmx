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

	// Tabs.
	mux.Get("/tabs", app.tabs)
	mux.Get("/tabs/{id}", app.tabs)

	// SSE.
	mux.Get("/events", app.events)
	mux.Get("/events-server", app.eventsServer)

	// Playing with refresh.
	mux.Get("/refresh", app.refresh)
	mux.Get("/refresh/server-time", app.refreshServerTime)

	// Colour swap.
	mux.Get("/colour-swap", app.colourSwap)
	mux.Get("/colour-swap/go", app.colourSwapGo)

	// Slide divs in / out.
	mux.Get("/slide", app.slide)
	mux.Get("/slide/go", app.slideGo)

	// Leaflet.
	mux.Get("/leaflet", app.leaflet)
	mux.Get("/leaflet/markers-in-bbox", app.leafletMarkersInBBox)
	mux.Get("/leaflet/marker-popup/{id}", app.leafletMarkerPopup)
	mux.Post("/leaflet/marker-choose", app.leafletMarkerChoose)

	// Confirm dialog.
	mux.Get("/confirm", app.confirm)
	mux.Get("/confirm/do-something", app.confirmDoSomething)

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
