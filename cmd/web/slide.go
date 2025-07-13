package main

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
)

var (
	ErrMissingSlidePage = errors.New("missing slide value")

	ErrInvalidDirection = errors.New("invalid direction")

	slidePages = []string{"one", "two", "three", "four", "five"}
)

func (app *Application) slide(w http.ResponseWriter, r *http.Request) {
	app.slidePage = slidePages[0]

	pageData := map[string]any{
		"Direction": "Next",
		"Page":      app.slidePage,
	}

	app.render(w, r, "slide", pageData, http.StatusOK)
}

func (app *Application) slideGo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	n := slices.Index(slidePages, app.slidePage)

	if n == -1 {
		app.serverError(
			w,
			r,
			fmt.Errorf("slices.BinarySearch(%s): %w", app.slidePage, ErrMissingSlidePage),
			http.StatusInternalServerError,
		)

		return
	}

	direction := r.FormValue("direction")
	if !slices.Contains([]string{"Next", "Prev"}, direction) {
		app.serverError(
			w,
			r,
			fmt.Errorf("slices.Contains(%s): %w", direction, ErrInvalidDirection),
			http.StatusInternalServerError,
		)

		return
	}

	if direction == "Next" {
		n++
	} else {
		n--
	}

	if n > len(slidePages)-1 {
		n = 0
	}
	if n < 0 {
		n = len(slidePages) - 1
	}

	app.slidePage = slidePages[n]

	pageData := map[string]any{
		"Page":      app.slidePage,
		"Direction": direction,
	}

	app.render(w, r, "slide-content", pageData, http.StatusOK)
}
