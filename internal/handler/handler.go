package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Serioga111/CutterService/internal/handler/dto"
	"github.com/Serioga111/CutterService/internal/repositorie"
	"github.com/Serioga111/CutterService/internal/service"
)

type Handler struct {
	storage   repositorie.Repositorie
	generator *service.Generator
}

func NewHandler(s repositorie.Repositorie) *Handler {
	return &Handler{
		storage:   s,
		generator: service.NewGenerator(s),
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /", h.createShortURL)
	mux.HandleFunc("GET /{short}", h.getOriginalURL)
}

func (h *Handler) createShortURL(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		sendJSONError(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortLink, err := h.generator.GenerateShortURL(req.URL)
	if err != nil {
		sendJSONError(w, "Failed to generate short link", http.StatusInternalServerError)
		return
	}

	savedShort, err := h.storage.Save(req.URL, shortLink)
	if err != nil {
		sendJSONError(w, "Failed to save link", http.StatusInternalServerError)
		return
	}

	resp := dto.CreateResponse{ShortURL: savedShort}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) getOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortLink := r.PathValue("short")

	if shortLink == "" {
		sendJSONError(w, "Short link is required", http.StatusBadRequest)
		return
	}

	originalURL, err := h.storage.Get(shortLink)
	if err != nil {
		sendJSONError(w, "Failed to get link", http.StatusInternalServerError)
		return
	}

	if originalURL == "" {
		sendJSONError(w, "Link not found", http.StatusNotFound)
		return
	}

	resp := dto.GetResponse{OriginalURL: originalURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(dto.ErrorResponse{Error: message})
}
