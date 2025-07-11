package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Agarwalsahil/TodoAPI/db"
	"github.com/Agarwalsahil/TodoAPI/models"
	"github.com/Agarwalsahil/TodoAPI/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dbTest, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	dbTest.AutoMigrate(&models.User{})
	db.DB = dbTest
	return dbTest
}

func TestLogin(t *testing.T) {
	dbTest := setupTestDB()

	hashed, _ := utils.HashPassword("secure@123")
	dbTest.Create(&models.User{Name: "Sahil", Email: "test@example.com", Password: hashed})

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", Login)

	body := `{
		"name": "Sahil",
		"email": "test@example.com",
		"password": "secure@123" 
	}`

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected statsu 200 but got %d", resp.Code)
	}
}
