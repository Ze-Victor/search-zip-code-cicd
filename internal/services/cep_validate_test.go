package services

import (
	"testing"
)

func TestValidateCEP(t *testing.T) {
	tests := []struct {
		name      string
		cep       string
		expectErr bool
	}{
		{"Valid CEP", "12345678", false},
		{"Invalid CEP Less than 8 digits", "1234567", true},
		{"Invalid CEP More than 8 digits", "123456789", true},
		{"Invalid CEP Contains letters", "1234abc6", true},
		{"Invalid CEP Empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCEP(tt.cep)
			if tt.expectErr && err == nil {
				t.Errorf("Expected error, but got none for test case: %s", tt.name)
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error %v for test case: %s", err, tt.name)
			}
		})
	}
}
