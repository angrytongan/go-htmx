package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	initialColour = "red"
)

type Application struct {
	tpl *template.Template

	colourSwapValue string
	slidePage       string
}

func newApplication() *Application {
	tpl := template.Must(template.ParseGlob("./templates/*.tmpl"))

	return &Application{
		tpl:             tpl,
		colourSwapValue: initialColour,
		slidePage:       "",
	}
}

func (app *Application) serverError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
	statusCode int,
) {
	log.Printf("%s %s: %v\n", r.Method, r.URL, err)
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func (app *Application) render(w http.ResponseWriter,
	r *http.Request,
	block string,
	pageData map[string]any,
	statusCode int,
) {
	var b bytes.Buffer

	if r.Header.Get("Hx-Request") == "" {
		block += "-page"
	}

	err := app.tpl.ExecuteTemplate(&b, block, pageData)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("app.tpl.ExecuteTemplate(%s): %w", block, err),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)

	_, err = w.Write(b.Bytes())
	if err != nil {
		app.serverError(w, r, fmt.Errorf("w.Write(): %w", err), http.StatusInternalServerError)

		return
	}
}
