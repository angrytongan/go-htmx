package main

import "net/http"

func (app *Application) root(w http.ResponseWriter, r *http.Request) {
	var pageData map[string]any

	app.render(w, r, "root", pageData, http.StatusOK)
}
