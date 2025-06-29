package main

import (
	"fmt"
	"net/http"

	"ghh/internal/words"
)

func (app *Application) wordsSearchRadio(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.serverError(w, r, fmt.Errorf("r.ParseForm(): %w", err), http.StatusInternalServerError)

		return
	}

	partial := r.FormValue("q")

	possibles, err := words.Search(partial)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("words.Search(%s): %w", partial, err),
			http.StatusInternalServerError,
		)

		return
	}

	pageData := map[string]any{
		"Name":    "word",
		"Results": possibles,
	}

	app.render(w, r, "word-search-results-radio", pageData, http.StatusOK)
}

func (app *Application) wordsSearchCheckbox(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.serverError(w, r, fmt.Errorf("r.ParseForm(): %w", err), http.StatusInternalServerError)

		return
	}

	partial := r.FormValue("q")

	possibles, err := words.Search(partial)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("words.Search(%s): %w", partial, err),
			http.StatusInternalServerError,
		)

		return
	}

	pageData := map[string]any{
		"Name":    "word",
		"Results": possibles,
	}

	app.render(w, r, "word-search-results-checkbox", pageData, http.StatusOK)
}
