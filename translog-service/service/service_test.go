package service_test

import (
	"testing"
	"translog-service/mocks"
	"translog-service/service"
	"github.com/golang/mock/gomock"
)

func TestValidateStatusTransition(t *testing.T) {
	svc := service.NewTranslogService(nil)

	err := svc.ValidateStatusTransition("SEARCHING", "COMPLETED")
	if err == nil {
		t.Errorf("Ekspektasi error untuk transisi SEARCHING -> COMPLETED, tapi sukses")
	}

	err = svc.ValidateStatusTransition("SEARCHING", "IN_PROGRESS")
	if err != nil {
		t.Errorf("Ekspektasi sukses untuk transisi SEARCHING -> IN_PROGRESS, tapi dapat error: %v", err)
	}
}

func TestCreateTransportOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTranslogRepository(ctrl)
	mockRepo.EXPECT().SaveOrder(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	svc := service.NewTranslogService(mockRepo)
	order, err := svc.CreateTransportOrder()

	if err != nil {
		t.Errorf("Tidak ekspektasi error, tapi dapat: %v", err)
	}
	if order.Status != "SEARCHING" {
		t.Errorf("Ekspektasi status SEARCHING, dapat: %s", order.Status)
	}
}

func TestFunctionalDBTranslogConnection(t *testing.T) {
	t.Skip("Skipping functional test unless explicitly needed (requires DB)")
}
