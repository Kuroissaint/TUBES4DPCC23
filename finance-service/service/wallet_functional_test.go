//go:build functional

package service_test

import (
	"testing"
)

func TestWallet_Functional(t *testing.T) {
	// Lewati functional test sementara sampai integrasi DB Staging selesai
	t.Skip("Skipping functional test wallet: DB Staging belum siap, akan diimplementasikan nanti.")
}
