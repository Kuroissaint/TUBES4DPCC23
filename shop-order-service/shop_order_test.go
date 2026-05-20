package main

import (
	"testing"
)

func TestAddToCartLogic(t *testing.T) {
	cart := ShoppingCart{
		Items: []string{},
	}

	cart.AddToCart("Mie Goreng Spesial")

	// Validasi panjang array
	if len(cart.Items) != 1 {
		t.Errorf("Ekspektasi item ada 1, tapi dapat: %d", len(cart.Items))
	}

	// Validasi isi array
	if cart.Items[0] != "Mie Goreng Spesial" {
		t.Errorf("Ekspektasi item 'Mie Goreng Spesial', tapi dapat: %s", cart.Items[0])
	}
}