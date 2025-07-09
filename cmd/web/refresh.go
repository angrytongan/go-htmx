package main

import (
	"math/rand"
	"net/http"
	"time"
)

func (app *Application) refresh(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "refresh", nil, http.StatusOK)
}

func (app *Application) refreshServerTime(w http.ResponseWriter, r *http.Request) {
	min := 0
	max := 10

	val := rand.Intn(max-min) + min
	if val >= 2 {
		pageData := map[string]any{
			"ServerTime": time.Now().String(),
			"Delay":      "1s",
		}

		app.render(w, r, "refresh-server-time", pageData, http.StatusOK)

		return
	}

	app.render(w, r, "refresh-server-time-paused", nil, http.StatusOK)
}
