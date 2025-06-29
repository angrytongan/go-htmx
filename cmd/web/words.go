package main

import (
	"fmt"
	"ghh/internal/words"
	"net/http"
)

func (app *Application) wordsSearchRadio(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	partial := r.FormValue("q")

	possibles, err := words.Search(partial)
	if err != nil {
		app.serverError(w, r, fmt.Errorf("words.Search(%s): %w", partial, err), http.StatusInternalServerError)

		return
	}

	pageData := map[string]any{
		"Name":    "word",
		"Results": possibles,
	}

	app.render(w, r, "word-search-results-radio", pageData, http.StatusOK)
}

func (app *Application) wordsSearchCheckbox(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	partial := r.FormValue("q")

	possibles, err := words.Search(partial)
	if err != nil {
		app.serverError(w, r, fmt.Errorf("words.Search(%s): %w", partial, err), http.StatusInternalServerError)

		return
	}

	pageData := map[string]any{
		"Name":    "word",
		"Results": possibles,
	}

	app.render(w, r, "word-search-results-checkbox", pageData, http.StatusOK)
}
