package main

import (
	"fmt"
	"net/http"

	"ghh/internal/news"
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

	pageData := map[string]any{
		"News": newsPartial,
		/*
			"Weather": weather.Partial(offset, limit),
			"Sport":   sport.Partial(offset, limit),
		*/
	}

	app.render(w, r, "dashboard", pageData, http.StatusOK)
}
