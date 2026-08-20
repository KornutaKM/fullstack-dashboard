package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KornutaKM/fullstack-dashboard/internal/services"
)

type ExchangeHandler struct {
	service *services.ExchangeService
}

func NewExchangeHandler(service *services.ExchangeService) *ExchangeHandler {
	return &ExchangeHandler{service: service}
}

func (h *ExchangeHandler) GetExchange(w http.ResponseWriter, r *http.Request) {
	base := r.URL.Query().Get("base")
	if base == "" {
		base = "USD"
	}

	result, err := h.service.GetExchange(r.Context(), base)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
