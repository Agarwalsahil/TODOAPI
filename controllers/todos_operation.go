package controllers

import (
	"fmt"
	"net/http"

	"github.com/Agarwalsahil/TodoAPI/db"
	"github.com/Agarwalsahil/TodoAPI/models"
	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var input models.Todo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo input"})
		return
	}

	createdBy := c.GetString("CreatedBy")
	if createdBy == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	var count int64
	db.DB.Model(&models.Todo{}).Where("created_by = ?", createdBy).Count(&count)

	todo := models.Todo{
		TodoID:      count + 1,
		Title:       input.Title,
		Description: input.Description,
	}

	if err := db.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Todo created succesfully", "todo": todo})
}

func GetTodos(c *gin.Context) {
	createdBy := c.GetString("CreatedBy")

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageNum := 1
	limitNum := 1
	fmt.Sscanf(page, "%d", &pageNum)
	fmt.Sscanf(limit, "%d", &limitNum)

	if pageNum < 1 {
		pageNum = 1
	}

	if limitNum < 1 {
		limitNum = 1
	}

	offset := (pageNum - 1)* limitNum

	var todos []models.Todo
	query := db.DB.Where("created_by = ?", createdBy)

	result := query.Offset(offset).Limit(limitNum).Find(&todos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func UpdateTodo(c *gin.Context){
	createdBy := c.GetString("CreatedBy")

	if createdBy == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	var todoID int64
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &todoID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	var todo models.Todo
	if err := db.DB.Where("todo_id = ? AND created_by = ?", todoID, createdBy).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo.Title = input.Title
	todo.Description = input.Description

	if err := db.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated", "todo": todo})
}

func DeleteTodo(c *gin.Context) {
	createdBy := c.GetString("CreatedBy")

	if createdBy == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	var todoID int64
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &todoID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input id"})
		return
	}

	var todo models.Todo
	if err := db.DB.Where("todo_id = ? AND created_by = ?", todoID, createdBy).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := db.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})

}
