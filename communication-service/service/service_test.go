package service_test

import (
	"context"
	"testing"
    
	"github.com/golang/mock/gomock"
	"github.com/kuroissaint/tubes2dpcc/communication-service/model"
	"github.com/kuroissaint/tubes2dpcc/communication-service/service"
	"github.com/kuroissaint/tubes2dpcc/communication-service/mocks" 
)

func TestProcessIncomingMessage_FailedEmptyOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockChatRepository(ctrl)
	svc := service.NewChatService(mockRepo)

	invalidMsg := model.ChatMessage{
		OrderID:     "", 
		SenderID:    "USR-123",
		MessageText: "Halo",
	}

	mockRepo.EXPECT().SaveMessage(gomock.Any(), gomock.Any()).Times(0)

	err := svc.ProcessIncomingMessage(context.Background(), invalidMsg)

	if err == nil || err.Error() != "order_id dan sender_id tidak boleh kosong" {
		t.Errorf("Ekspektasi error validasi, tapi mendapat: %v", err)
	}
}

func TestProcessIncomingMessage_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockChatRepository(ctrl)
	svc := service.NewChatService(mockRepo)

	validMsg := model.ChatMessage{
		OrderID:     "ORD-123",
		SenderID:    "USR-123",
		MessageText: "Sesuai aplikasi ya",
	}

	mockRepo.EXPECT().SaveMessage(gomock.Any(), validMsg).Return(nil).Times(1)

	err := svc.ProcessIncomingMessage(context.Background(), validMsg)

	if err != nil {
		t.Errorf("Ekspektasi sukses, tapi ada error: %v", err)
	}
}
