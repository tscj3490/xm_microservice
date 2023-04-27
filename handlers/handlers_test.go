package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const BASE_URL = "http://localhost:9090"

func TestGetCompanyHandler(t *testing.T) {
	// Load environment variables from the .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// create a test server with the GetCompanyHandler function
	handler := GetCompanyHandler()
	server := httptest.NewServer(handler)
	defer server.Close()
	fmt.Println("Server URL:", server.URL)

	// send a GET request to retrieve a company with ID 1
	req, err := http.NewRequest("GET", BASE_URL+"/companies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the auth
	req.SetBasicAuth(os.Getenv("USER_NAME"), os.Getenv("USER_PASS"))
	// req.Header.Set("Authorization", tt.header)

	client := &http.Client{}
	fmt.Println("Request URL:", req.URL)
	res, err := client.Do(req)

	if err != nil {
		t.Fatal("Cannot create http request")
	}
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCreateCompanyHandler(t *testing.T) {

	// Load environment variables from the .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Create a new request with some JSON data
	reqBody := []byte(`{"id": "5", "name": "facebook", "description": "social company", "employees": 36000, "registered": true, "type": "Corporations"}`)
	req, err := http.NewRequest("POST", BASE_URL+"/companies", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set the auth
	req.SetBasicAuth(os.Getenv("USER_NAME"), os.Getenv("USER_PASS"))

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal("Cannot create http request")
	}
	defer res.Body.Close()

	assert.Equal(t, http.StatusCreated, res.StatusCode)
}
