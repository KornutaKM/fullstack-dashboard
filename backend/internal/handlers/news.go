package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KornutaKM/fullstack-dashboard/internal/services"
)

type NewsHandler struct {
	service *services.NewsService
}

func NewNewsHandler(service *services.NewsService) *NewsHandler {
	return &NewsHandler{service: service}
}

func (h *NewsHandler) GetNews(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	if country == "" {
		country = "us"
	}

	category := r.URL.Query().Get("category")
	if category == "" {
		category = "general"
	}

	news, err := h.service.GetTopHeadlines(r.Context(), country, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}
