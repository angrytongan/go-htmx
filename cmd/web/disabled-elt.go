package main

import (
	"net/http"
	"time"
)

func (app *Application) disabledElt(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "disabled-elt", nil, http.StatusOK)
}

func (app *Application) disabledEltReq(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)

	pageData := map[string]any{
		"ServerTime": time.Now().String(),
	}

	app.render(w, r, "disabled-elt-result", pageData, http.StatusOK)
}
