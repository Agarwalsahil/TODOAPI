package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegister(t *testing.T) {
	setupTestDB()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", Register)

	body := `{
		"name":"Sahil",
		"email": "test@example.com",
		"password": "test123"
	}`

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.Code)
	}
}
