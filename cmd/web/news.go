package main

import (
	"fmt"
	"net/http"
	"strconv"

	"ghh/internal/news"
)

func (app *Application) newsPartial(w http.ResponseWriter, r *http.Request) {
	offsetStr := r.PathValue("offset")
	limitStr := r.PathValue("limit")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("strconv.Atoi(offset, %s): %w", offsetStr, err),
			http.StatusBadRequest,
		)

		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("strconv.Atoi(limit, %s): %w", limitStr, err),
			http.StatusBadRequest,
		)

		return
	}

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
		"Items":    newsPartial.Items,
		"HasPrev":  newsPartial.HasPrev,
		"HasNext":  newsPartial.HasNext,
		"NextPage": newsPartial.NextPage,
		"PrevPage": newsPartial.PrevPage,
		"Limit":    newsPartial.Limit,
	}

	app.render(w, r, "news", pageData, http.StatusOK)
}
