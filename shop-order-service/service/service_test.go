package service_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	
	"shop-order-service/mocks"
	"shop-order-service/model"
	"shop-order-service/service"

	"github.com/golang/mock/gomock"
)

// Bapak buatkan MockTransport untuk mencegat (intercept) HTTP request bawaan Golang
type MockTransport struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}

// Test logika bisnis AddToCart (Unit Test murni)
func TestAddToCartLogic(t *testing.T) {
	cart := model.ShoppingCart{Items: []string{}}
	cart.AddToCart("Mie Goreng Spesial")

	if len(cart.Items) != 1 || cart.Items[0] != "Mie Goreng Spesial" {
		t.Errorf("Add To Cart logic failed: ekspektasi [Mie Goreng Spesial], dapat %v", cart.Items)
	}
}

// Test integrasi service dengan repository menggunakan Mock jaringan dan database
func TestCreateShoppingOrder(t *testing.T) {
	// PENCEGAHAN ERROR NO SUCH HOST:
	// Kita simpan transport asli, dan ganti Transport bawaan klien HTTP Golang dengan buatan kita.
	// Sehingga fungsi http.Post ke "translog-service" akan langsung sukses tanpa butuh koneksi.
	originalTransport := http.DefaultClient.Transport
	defer func() {
		http.DefaultClient.Transport = originalTransport // Kembalikan ke semula setelah tes selesai
	}()

	http.DefaultClient.Transport = &MockTransport{
		RoundTripFunc: func(req *http.Request) (*http.Response, error) {
			// Memalsukan respons sukses (200 OK) dari translog-service
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`{"status": "success", "message": "Order accepted"}`)),
				Header:     make(http.Header),
			}, nil
		},
	}

	// 1. Inisialisasi Gomock Controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 2. Buat mock dari repository
	mockRepo := mocks.NewMockShopOrderRepository(ctrl)

	// 3. Set Ekspektasi: SaveCart
	mockRepo.EXPECT().SaveCart(
		gomock.Any(), // orderID
		gomock.Any(), // userID
		gomock.Any(), // merchantID
		gomock.Any(), // items
		gomock.Any(), // status
	).Return(nil).Times(1)

	// 4. Injeksi mock ke service
	svc := service.NewShopOrderService(mockRepo)

	// 5. Jalankan fungsi
	// Bapak perbaiki inputnya agar berisi Kopi & Roti, sesuai dengan ekspektasi Assertions kamu
	inputCart := &model.ShoppingCart{
		Items: []string{"Kopi", "Roti"},
	}
	cart, err := svc.CreateShoppingOrder(inputCart)

	// 6. Assertions (pengecekan hasil)
	if err != nil {
		t.Errorf("Tidak ekspektasi error, tapi dapat: %v", err)
	}
	
	if cart == nil {
		t.Fatal("Ekspektasi cart tidak nil")
	}

	if len(cart.Items) != 2 {
		t.Errorf("Ekspektasi 2 item (Kopi & Roti), tapi dapat: %d", len(cart.Items))
	}
}

// Test fungsional database
func TestFunctionalDBShoppingConnection(t *testing.T) {
	t.Skip("Skipping functional test unless explicitly needed (requires DB)")
}