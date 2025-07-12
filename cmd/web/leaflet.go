package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

type Marker struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var (
	townMarkers = []Marker{
		{
			ID:        0,
			Title:     "Mooloolah Valley",
			Latitude:  -26.765767,
			Longitude: 152.963076,
		},
		{
			ID:        1,
			Title:     "Landsborough",
			Latitude:  -26.811662,
			Longitude: 152.965565,
		},
		{
			ID:        2,
			Title:     "Maroochydore",
			Latitude:  -26.655935,
			Longitude: 153.094954,
		},
		{
			ID:        3,
			Title:     "Coolum",
			Latitude:  -26.528951,
			Longitude: 153.091736,
		},
		{
			ID:        4,
			Title:     "Caloundra",
			Latitude:  -26.798294,
			Longitude: 153.142333,
		},
	}
)

func (app *Application) leaflet(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "leaflet", nil, http.StatusOK)
}

func (app *Application) leafletMarkersInBBox(w http.ResponseWriter, r *http.Request) {
	markers := []Marker{}

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, r, fmt.Errorf("r.ParseForm(): %w", err), http.StatusInternalServerError)

		return
	}

	bbox := r.FormValue("bbox")
	tokens := strings.Split(bbox, ",")
	swLng, _ := strconv.ParseFloat(tokens[0], 32)
	swLat, _ := strconv.ParseFloat(tokens[1], 32)
	neLng, _ := strconv.ParseFloat(tokens[2], 32)
	neLat, _ := strconv.ParseFloat(tokens[3], 32)

	for _, m := range townMarkers {
		if swLat <= m.Latitude && m.Latitude <= neLat && swLng <= m.Longitude && m.Longitude <= neLng {
			markers = append(markers, m)
		}
	}

	fmt.Println(markers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(markers)
}

func (app *Application) leafletChooseMarker(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")

	markerID, err := strconv.Atoi(id)
	if err != nil {
		app.serverError(w, r, fmt.Errorf("strconv.Atoi(%s): %w", id, err), http.StatusInternalServerError)

		return
	}

	idx := slices.IndexFunc(townMarkers, func(m Marker) bool {
		return m.ID == markerID
	})
	if idx == -1 {
		app.serverError(w, r, fmt.Errorf("slices.IndexFunc(%s): %w", id, err), http.StatusBadRequest)

		return
	}

	pageData := map[string]any{
		"Chosen": townMarkers[idx],
	}

	app.render(w, r, "leaflet-chosen-marker", pageData, http.StatusOK)
}
