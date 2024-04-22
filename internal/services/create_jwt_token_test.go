package services

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		password    string
		expectedErr bool
	}{
		{"ValidCredentials", "test@example.com", "secretpassword", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := CreateToken(tt.email, tt.password)
			if tt.expectedErr && (err == nil || token != "") {
				t.Errorf("Expected error for creating token")
			} else {
				if err != nil {
					t.Errorf("Error creating token: %v", err)
				}
				if token == "" {
					t.Errorf("Token should not be empty")
				}
			}
		})
	}
}
