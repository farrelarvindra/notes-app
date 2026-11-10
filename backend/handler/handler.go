package handler

import (
	"encoding/json"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlerMain)
	mux.HandleFunc("GET /health", handlerHealth)
	return mux
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_ = writeJSON(w, map[string]string{
		"message": "welcome to notes app api",
	})
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	_ = writeJSON(w, map[string]string{
		"status": "ok",
	})
}

// writeJSON sets the Content-Type header and encodes the given map into JSON.
func writeJSON(w http.ResponseWriter, data map[string]string) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
