package main

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
)

var (
	ErrMissingSlidePage = errors.New("missing slide value")

	slidePages = []string{"one", "two", "three", "four", "five"}
)

func (app *Application) slide(w http.ResponseWriter, r *http.Request) {
	app.slidePage = slidePages[0]

	app.render(w, r, "slide", nil, http.StatusOK)
}

func (app *Application) slideGo(w http.ResponseWriter, r *http.Request) {
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

	pageData := map[string]any{
		"Page": app.slidePage,
	}

	n++
	if n > len(slidePages)-1 {
		n = 0
	}
	app.slidePage = slidePages[n]

	app.render(w, r, "slide", pageData, http.StatusOK)
}
