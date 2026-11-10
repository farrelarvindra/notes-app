package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/farrelarvindra/notes-app/model"
)

// Handler holds dependencies for HTTP handlers.
type Handler struct {
	DB *gorm.DB
}

// New creates a new Handler with GORM DB.
func New(db *gorm.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

// Routes sets up all API routes and middleware.
func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	// Endpoints
	mux.HandleFunc("GET /{$}", h.handlerMain)
	mux.HandleFunc("GET /api/{$}", h.handlerMain)
	mux.HandleFunc("GET /api/health", h.handlerHealth)
	mux.HandleFunc("GET /api/notes", h.handlerGetNotes)
	mux.HandleFunc("POST /api/notes", h.handlerCreateNote)
	mux.HandleFunc("PUT /api/notes/{id}", h.handlerUpdateNote)
	mux.HandleFunc("DELETE /api/notes/{id}", h.handlerDeleteNote)

	// Wrap with CORS middleware
	return enableCORS(mux)
}

func (h *Handler) handlerMain(w http.ResponseWriter, r *http.Request) {
	_ = writeJSON(w, http.StatusOK, map[string]string{
		"message": "welcome to notes app api",
	})
}

func (h *Handler) handlerHealth(w http.ResponseWriter, r *http.Request) {
	_ = writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h *Handler) handlerGetNotes(w http.ResponseWriter, r *http.Request) {
	var notes []model.Note
	if err := h.DB.Order("created_at desc").Find(&notes).Error; err != nil {
		_ = writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to retrieve notes",
		})
		return
	}

	// Return empty array [] rather than null if no records found
	if notes == nil {
		notes = []model.Note{}
	}

	_ = writeJSON(w, http.StatusOK, notes)
}

func (h *Handler) handlerCreateNote(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if strings.TrimSpace(req.Title) == "" {
		_ = writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "title is required",
		})
		return
	}

	note := model.Note{
		Title:   strings.TrimSpace(req.Title),
		Content: req.Content,
	}

	if err := h.DB.Create(&note).Error; err != nil {
		_ = writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to create note",
		})
		return
	}

	_ = writeJSON(w, http.StatusCreated, note)
}

func (h *Handler) handlerUpdateNote(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		_ = writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid note id",
		})
		return
	}

	var note model.Note
	if err := h.DB.First(&note, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = writeJSON(w, http.StatusNotFound, map[string]string{
				"error": "note not found",
			})
			return
		}
		_ = writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to find note",
		})
		return
	}

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if strings.TrimSpace(req.Title) == "" {
		_ = writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "title is required",
		})
		return
	}

	note.Title = strings.TrimSpace(req.Title)
	note.Content = req.Content

	if err := h.DB.Save(&note).Error; err != nil {
		_ = writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to update note",
		})
		return
	}

	_ = writeJSON(w, http.StatusOK, note)
}

func (h *Handler) handlerDeleteNote(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		_ = writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid note id",
		})
		return
	}

	var note model.Note
	if err := h.DB.First(&note, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = writeJSON(w, http.StatusNotFound, map[string]string{
				"error": "note not found",
			})
			return
		}
		_ = writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to find note",
		})
		return
	}

	if err := h.DB.Delete(&note).Error; err != nil {
		_ = writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to delete note",
		})
		return
	}

	_ = writeJSON(w, http.StatusOK, map[string]string{
		"message": "note deleted successfully",
	})
}

// writeJSON sets the Content-Type header, status code, and encodes data into JSON.
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
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
