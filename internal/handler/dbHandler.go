package handler

import (
	"net/http"

	storage "github.com/Serioga111/CutterService/internal/repositorie"
)

type Handler struct {
	storage storage.Storage
}

func NewHandler(s storage.Storage) *Handler {
	return &Handler{
		storage: s,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /{link}", h.createShortLink)
	mux.HandleFunc("GET /{short}", h.getOriginalLink)
}

func (h *Handler) createShortLink(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("shorten"))
}

func (h *Handler) getOriginalLink(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("original"))
}
