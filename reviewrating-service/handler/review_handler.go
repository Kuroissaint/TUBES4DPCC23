package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/model"
	"github.com/kuroissaint/tubes2dpcc/reviewrating-service/service"
)

type ReviewHandler struct {
	reviewService *service.ReviewService
}

func NewReviewHandler(rs *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{reviewService: rs}
}

func (h *ReviewHandler) SubmitReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req model.Review
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	err := h.reviewService.SubmitReview(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "review submitted successfully"})
}
