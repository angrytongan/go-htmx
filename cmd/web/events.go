package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (app *Application) events(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "events", nil, http.StatusOK)
}

func (app *Application) eventsServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	clientGone := r.Context().Done()

	rc := http.NewResponseController(w)
	t := time.NewTicker(time.Second)
	t2 := time.NewTicker(2 * time.Second)
	defer t.Stop()
	defer t2.Stop()

	for {
		select {
		case <-clientGone:
			log.Println("client disconnected")

			return

		case <-t.C:
			_, err := fmt.Fprintf(w, "event: Event1\ndata: <div>time is %s</div>\n\n", time.Now())
			if err != nil {
				log.Println("select: %w", err)

				return
			}

			if err := rc.Flush(); err != nil {
				log.Println("rc.Flush(): %w", err)

				return
			}

		case <-t2.C:
			_, err := fmt.Fprintf(w, "event: Event2\ndata: <div>time is %s</div>\n\n", time.Now())
			if err != nil {
				log.Println("select: %w", err)

				return
			}

			if err := rc.Flush(); err != nil {
				log.Println("rc.Flush(): %w", err)

				return
			}
		}
	}
}
