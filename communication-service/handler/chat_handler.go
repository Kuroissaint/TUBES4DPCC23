package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kuroissaint/tubes2dpcc/communication-service/model"
	"github.com/kuroissaint/tubes2dpcc/communication-service/service"
)

type ChatHandler struct {
	chatService *service.ChatService
}

func NewChatHandler(cs *service.ChatService) *ChatHandler {
	return &ChatHandler{chatService: cs}
}

func (h *ChatHandler) ProcessMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req model.ChatMessage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	err := h.chatService.ProcessIncomingMessage(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "message processed successfully"})
}
