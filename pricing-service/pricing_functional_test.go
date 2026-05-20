//go:build functional

package main

import (
	"testing"
)

func TestCalculateFinalPrice_Functional(t *testing.T) {
	// Lewati functional test sementara sampai integrasi DB Staging selesai
	t.Skip("Skipping functional test: DB Staging belum siap, akan diimplementasikan nanti.")
}