package server

import (
	"net/http"
)

func NewMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", HealthHandler)

	return mux
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"status\":\"OK\"}"))
}
