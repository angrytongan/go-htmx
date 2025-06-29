package main

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	numGraphs = 3
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

func makeAllGraphs(n int) ([]GraphThing, error) {
	graphs := []GraphThing{}

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
			return graphs, fmt.Errorf("bar.Render: %w", err)
		}

		snippet := bar.RenderSnippet()

		graphs = append(graphs, GraphThing{
			Element: template.HTML(snippet.Element),
			Script:  template.HTML(snippet.Script),
		})
	}

	return graphs, nil
}

func (app *Application) graphs(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	graphs, err := makeAllGraphs(numGraphs)
	renderDuration := time.Since(start)

	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("makeAllGraphs(%d): %w", numGraphs, err),
			http.StatusInternalServerError,
		)

		return
	}

	pageData := map[string]any{
		"Graphs":        graphs,
		"Servertime":    time.Now().String(),
		"RenderingTime": renderDuration,
	}

	// We use the same handler for normal loads and htmx loads. Difference is
	// what template we use on the rendering.
	block := "graphs"
	if r.Header.Get("Hx-Request") == "true" {
		block = "all-graphs"
	}

	app.render(w, r, block, pageData, http.StatusOK)
}
