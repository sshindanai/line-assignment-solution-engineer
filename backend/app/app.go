package app

import (
	"net/http"

	"github.com/r3labs/sse"
)

var (
	server *sse.Server
)

func StartApplication() {
	server = sse.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/events", server.HTTPHandler)
	http.ListenAndServe(":8080", mux)
}
