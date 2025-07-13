package main

import "net/http"

func (app *Application) colourSwap(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "colour-swap", nil, http.StatusOK)
}

func (app *Application) colourSwapGo(w http.ResponseWriter, r *http.Request) {
	if app.colourSwapValue == "red" {
		app.colourSwapValue = "green"
	} else {
		app.colourSwapValue = "red"
	}

	pageData := map[string]any{
		"Colour": app.colourSwapValue,
	}

	app.render(w, r, "color-swap-go", pageData, http.StatusOK)
}
