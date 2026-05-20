package main

import (
	"testing"
)

func TestValidateStatusTransition(t *testing.T) {
	// Skenario 1: Transisi tidak valid
	err := ValidateStatusTransition("SEARCHING", "COMPLETED")
	if err == nil {
		t.Errorf("Ekspektasi error untuk transisi SEARCHING -> COMPLETED, tapi sukses")
	}

	// Skenario 2: Transisi valid
	err = ValidateStatusTransition("SEARCHING", "IN_PROGRESS")
	if err != nil {
		t.Errorf("Ekspektasi sukses untuk transisi SEARCHING -> IN_PROGRESS, tapi dapat error: %v", err)
	}
}