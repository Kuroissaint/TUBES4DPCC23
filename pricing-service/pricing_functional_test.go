//go:build functional

package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculateFinalPrice_Functional(t *testing.T) {
	// Di sini nanti tempat ngetest langsung nembak DB lokal/staging
	// cth: db.Connect()...
	
	// Sengaja di-failed kan karena fitur belum siap
	assert.Fail(t, "Functional test sengaja gagal: Fitur belum siap di-test ke DB Staging.")
}