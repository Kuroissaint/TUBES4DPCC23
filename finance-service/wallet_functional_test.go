//go:build functional

package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWallet_Functional(t *testing.T) {
	// Tempat ngetest transaksi langsung ke DB Wallet Staging/Lokal
	assert.Fail(t, "Functional test wallet sengaja gagal: Fitur belum siap di-test ke DB Staging.")
}