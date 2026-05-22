package shoporder_test

import (
	"testing"
	
	shoporder "shop-order-service"
	"shop-order-service/mocks"

	"github.com/golang/mock/gomock"
)

// Test logika bisnis AddToCart (Unit Test murni)
func TestAddToCartLogic(t *testing.T) {
	cart := shoporder.ShoppingCart{Items: []string{}}
	cart.AddToCart("Mie Goreng Spesial")

	if len(cart.Items) != 1 || cart.Items[0] != "Mie Goreng Spesial" {
		t.Errorf("Add To Cart logic failed: ekspektasi [Mie Goreng Spesial], dapat %v", cart.Items)
	}
}

// Test integrasi service dengan repository menggunakan Mock
func TestCreateShoppingOrder(t *testing.T) {
	// 1. Inisialisasi Gomock Controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 2. Buat mock dari repository
	mockRepo := mocks.NewMockShopOrderRepository(ctrl)

	// 3. Set Ekspektasi: SaveCart akan dipanggil dengan parameter apa saja (gomock.Any())
	// Kita kembalikan 'nil' karena kita asumsikan simpan ke database berhasil
	mockRepo.EXPECT().SaveCart(
		gomock.Any(), // orderID
		gomock.Any(), // userID
		gomock.Any(), // merchantID
		gomock.Any(), // items
		gomock.Any(), // status
	).Return(nil).Times(1)

	// 4. Injeksi mock ke service
	svc := shoporder.NewShopOrderService(mockRepo)

	// 5. Jalankan fungsi
	cart, err := svc.CreateShoppingOrder()

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
	// Skip jika tidak di-build dengan tag functional
	// Jika menggunakan go test biasa, test ini mungkin gagal jika DB belum siap.
	// Namun agar struktur konsisten, kita biarkan logicnya di sini.
	t.Skip("Skipping functional test unless explicitly needed (requires DB)")
	/*
	connStr := "user=postgres password=password dbname=db_shopping_order sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Gagal inisialisasi koneksi: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Errorf("Functional Test Failed: Database db_shopping_order belum siap atau belum di-setup: %v", err)
	}
	*/
}
