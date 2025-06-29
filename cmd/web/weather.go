package main

import (
	"fmt"
	"net/http"
	"strconv"

	"ghh/internal/weather"
)

func (app *Application) weatherPartial(w http.ResponseWriter, r *http.Request) {
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
		"Items":    weatherPartial.Items,
		"HasPrev":  weatherPartial.HasPrev,
		"HasNext":  weatherPartial.HasNext,
		"NextPage": weatherPartial.NextPage,
		"PrevPage": weatherPartial.PrevPage,
		"Limit":    weatherPartial.Limit,
	}

	app.render(w, r, "weather", pageData, http.StatusOK)
}
