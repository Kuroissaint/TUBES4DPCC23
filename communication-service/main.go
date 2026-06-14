package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kuroissaint/tubes2dpcc/communication-service/handler"
	"github.com/kuroissaint/tubes2dpcc/communication-service/model"
	"github.com/kuroissaint/tubes2dpcc/communication-service/repository"
	"github.com/kuroissaint/tubes2dpcc/communication-service/service"
)

type dummyChatRepo struct{}

func (r *dummyChatRepo) SaveMessage(ctx context.Context, msg model.ChatMessage) error {
	return nil
}

func main() {
	var repo repository.ChatRepository = &dummyChatRepo{}
	svc := service.NewChatService(repo)
	hdl := handler.NewChatHandler(svc)

	http.HandleFunc("/api/chat/process", hdl.ProcessMessageHandler)

	fmt.Println("Communication Service running on :8082")
	http.ListenAndServe(":8082", nil)
}
