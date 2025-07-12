package main

import (
	"net/http"
)

func (app *Application) tabs(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Hx-Request") != "true" {
		app.render(w, r, "tabs", nil, http.StatusOK)

		return
	}

	content := map[string]string{
		"1": "this is content 1",
		"2": "this is content 2",
		"3": "this is content 3",
	}

	id := r.PathValue("id")

	pageData := map[string]any{
		"ID":        id,
		"Available": []string{"1", "2", "3"},
		"Content":   content[id],
	}

	app.render(w, r, "tabs", pageData, http.StatusOK)
}
