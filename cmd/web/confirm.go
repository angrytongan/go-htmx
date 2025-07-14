package main

import (
	"math/rand/v2"
	"net/http"
)

func (app *Application) confirm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "confirm", nil, http.StatusOK)
}

func (app *Application) confirmDoSomething(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Number": rand.IntN(5),
	}

	app.render(w, r, "confirm-do-something", pageData, http.StatusOK)
}
