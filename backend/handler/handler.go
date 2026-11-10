package handler

import (
	"encoding/json"
	"net/http"
)

// Routes sets up all API routes and middleware.
func Routes() http.Handler {
	mux := http.NewServeMux()

	// Endpoints
	mux.HandleFunc("GET /{$}", handlerMain)
	mux.HandleFunc("GET /api/{$}", handlerMain)
	mux.HandleFunc("GET /api/health", handlerHealth)

	// Wrap with CORS middleware
	return enableCORS(mux)
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	_ = writeJSON(w, map[string]string{
		"message": "welcome to notes app api",
	})
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	_ = writeJSON(w, map[string]string{
		"status": "ok",
	})
}

// writeJSON sets the Content-Type header and encodes the given data into JSON.
func writeJSON(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

// enableCORS handles Cross-Origin Resource Sharing for frontend access.
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
