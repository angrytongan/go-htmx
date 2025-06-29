package main

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type GraphThing struct {
	Element template.HTML
	Script  template.HTML
}

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for range 7 {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}

	return items
}

func (app *Application) graphs(w http.ResponseWriter, r *http.Request) {
	graphs := []GraphThing{}
	n := 3

	for range n {
		bar := charts.NewBar()
		b := bytes.Buffer{}

		bar.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{
				Height: "200px",
			}),
		)

		// Put data into instance
		bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
			AddSeries("Category A", generateBarItems()).
			AddSeries("Category B", generateBarItems())

		if err := bar.Render(&b); err != nil {
			app.serverError(w, r, fmt.Errorf("bar.Render: %w", err), http.StatusInternalServerError)

			return
		}

		snippet := bar.RenderSnippet()

		graphs = append(graphs, GraphThing{
			Element: template.HTML(snippet.Element),
			Script:  template.HTML(snippet.Script),
		})
	}

	pageData := map[string]any{
		"Graphs": graphs,
	}

	app.render(w, r, "graphs", pageData, http.StatusOK)
}
