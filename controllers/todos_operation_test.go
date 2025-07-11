package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Agarwalsahil/TodoAPI/db"
	"github.com/Agarwalsahil/TodoAPI/middleware"
	"github.com/Agarwalsahil/TodoAPI/models"
	"github.com/Agarwalsahil/TodoAPI/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTodoTest(t *testing.T) (*gin.Engine, string) {
	dbTest, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test db: %v", err)
	}

	// Auto-migrate models
	dbTest.AutoMigrate(&models.Todo{})
	db.DB = dbTest

	email := "testuser@example.com"
	token, _ := utils.GenerateToken(email)

	r := gin.Default()
	protected := r.Group("/todos")
	protected.Use(middleware.AuthMiddleware)

	protected.POST("", CreateTodo)
	protected.GET("", GetTodos)
	protected.PUT("/:id", UpdateTodo)
	protected.DELETE("/:id", DeleteTodo)

	return r, token
}

func TestCreateTodo(t *testing.T) {
	router, token := setupTodoTest(t)

	todo := map[string]string{
		"title":       "Test Task",
		"description": "This is a test task",
	}
	body, _ := json.Marshal(todo)

	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestGetTodos(t *testing.T) {
	router, token := setupTodoTest(t)

	// First create a todo
	db.DB.Create(&models.Todo{
		Title:       "Sample",
		Description: "Desc",
		CreatedBy:   "testuser@example.com",
	})

	req, _ := http.NewRequest("GET", "/todos", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateTodo(t *testing.T) {
	router, token := setupTodoTest(t)

	// Insert a todo
	todo := models.Todo{
		Title:       "Old Title",
		Description: "Old Description",
		CreatedBy:   "testuser@example.com",
	}
	db.DB.Create(&todo)

	update := map[string]string{
		"title":       "Updated Title",
		"description": "Updated Description",
	}
	body, _ := json.Marshal(update)

	req, _ := http.NewRequest("PUT", "/todos/"+fmt.Sprint(todo.TodoID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteTodo(t *testing.T) {
	router, token := setupTodoTest(t)

	todo := models.Todo{
		Title:       "Delete Me",
		Description: "Delete this todo",
		CreatedBy:   "testuser@example.com",
	}
	db.DB.Create(&todo)

	req, _ := http.NewRequest("DELETE", "/todos/"+fmt.Sprint(todo.TodoID), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
