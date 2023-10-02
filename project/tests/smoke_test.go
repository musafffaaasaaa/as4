package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSmoke(t *testing.T) {
	// Define the base URL of your website
	baseURL := "http://localhost:8080"

	// Send a GET request to the home page
	resp, err := http.Get(baseURL)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 (OK)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}

	// You can add more checks here to validate other aspects of your website

	fmt.Println("Smoke Test Passed!")
}
