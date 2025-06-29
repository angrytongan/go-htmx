package main

import (
	"fmt"
	"net/http"

	"ghh/internal/news"
	"ghh/internal/weather"
)

func (app *Application) dashboard(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 4

	newsPartial, err := news.Partial(offset, limit)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("news.Partial(%d, %d): %w", offset, limit, err),
			http.StatusInternalServerError,
		)

		return
	}

	weatherPartial, err := weather.Partial(offset, limit)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("weather.Partial(%d, %d): %w", offset, limit, err),
			http.StatusInternalServerError,
		)

		return
	}

	pageData := map[string]any{
		"News":    newsPartial,
		"Weather": weatherPartial,
	}

	app.render(w, r, "dashboard", pageData, http.StatusOK)
}
