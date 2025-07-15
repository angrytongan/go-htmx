package main

import "github.com/go-chi/chi/v5"

func (app *Application) setRoutes(mux *chi.Mux) {
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

	// disabled-elt.
	mux.Get("/disabled-elt", app.disabledElt)
	mux.Get("/disabled-elt/req", app.disabledEltReq)
}
