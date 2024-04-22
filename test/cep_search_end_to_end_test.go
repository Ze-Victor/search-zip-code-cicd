package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Ze-Victor/search-zip-code/config"
	pkg "github.com/Ze-Victor/search-zip-code/internal/pkg/authorization"
)

func TestSearchCEPEndToEnd(t *testing.T) {
	testCases := []struct {
		Name           string
		CEP            string
		ExpectedStatus int
	}{
		{"Valid CEP", "02010010", http.StatusOK},
		{"Not Found CEP", "00000000", http.StatusNotFound},
		{"CEP with More Digits CEP", "123456789", http.StatusBadRequest},
		{"CEP with Less Digits CEP", "1234567", http.StatusBadRequest},
		{"CEP with Digits Non Numeric", "123abc45", http.StatusBadRequest},
		{"Request with JSON malformad", "{ 12345678 }", http.StatusBadRequest},
		{"Invalid Token", "12345678", http.StatusUnauthorized},
	}

	requestBody := map[string]string{
		"email":    "seu_email@example.com",
		"password": "sua_senha_secreta",
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	authToken, err := getAuthToken(t, bodyBytes)
	if err != nil {
		t.Fatalf("failed to get auth token: %v", err)
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			req, err := http.NewRequest("GET", fmt.Sprintf("%s/cep/%s", config.Base_Path, tc.CEP), nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			if tc.Name != "Invalid Token" {
				req.Header.Set("Authorization", authToken)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("failed to make request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.ExpectedStatus {
				t.Errorf("expected status %v for test '%s'; got %v", tc.ExpectedStatus, tc.Name, resp.Status)
			}
		})
	}
}

func getAuthToken(t *testing.T, requestBody []byte) (string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth", config.Base_Path), bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get auth token: %s", resp.Status)
	}

	var authResponse pkg.TokenJWT
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return "", err
	}
	return authResponse.Token, nil
}
