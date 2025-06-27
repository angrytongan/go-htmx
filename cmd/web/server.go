package main

import (
	"net/http"
	"strconv"
)

func newServer(port int, handler http.Handler) *http.Server {
	addr := ":" + strconv.Itoa(port)

	return &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 0,
	}
}
