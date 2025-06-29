package main

import "net/http"

func (app *Application) comboboxes(w http.ResponseWriter, r *http.Request) {
	var pageData map[string]any

	app.render(w, r, "comboboxes", pageData, http.StatusOK)
}
