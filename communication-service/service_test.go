package communication_test

import (
	"context"
	"testing"
    
	"github.com/golang/mock/gomock"
	"github.com/kuroissaint/tubes2dpcc/communication-service"
	"github.com/kuroissaint/tubes2dpcc/communication-service/mocks" 
)

func TestProcessIncomingMessage_FailedEmptyOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockChatRepository(ctrl)
	service := communication.NewChatService(mockRepo)

	invalidMsg := communication.ChatMessage{
		OrderID:     "", // Kosong, harus gagal
		SenderID:    "USR-123",
		MessageText: "Halo",
	}

	// Ekspektasi: Database tidak boleh dipanggil karena gagal validasi
	mockRepo.EXPECT().SaveMessage(gomock.Any(), gomock.Any()).Times(0)

	err := service.ProcessIncomingMessage(context.Background(), invalidMsg)

	if err == nil || err.Error() != "order_id dan sender_id tidak boleh kosong" {
		t.Errorf("Ekspektasi error validasi, tapi mendapat: %v", err)
	}
}

func TestProcessIncomingMessage_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockChatRepository(ctrl)
	service := communication.NewChatService(mockRepo)

	validMsg := communication.ChatMessage{
		OrderID:     "ORD-123",
		SenderID:    "USR-123",
		MessageText: "Sesuai aplikasi ya",
	}

	// Ekspektasi: Database dipanggil 1 kali dan mengembalikan sukses (nil)
	mockRepo.EXPECT().SaveMessage(gomock.Any(), validMsg).Return(nil).Times(1)

	err := service.ProcessIncomingMessage(context.Background(), validMsg)

	if err != nil {
		t.Errorf("Ekspektasi sukses, tapi ada error: %v", err)
	}
}