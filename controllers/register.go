package controllers

import (
	"net/http"

	"github.com/Agarwalsahil/TodoAPI/db"
	"github.com/Agarwalsahil/TodoAPI/models"
	"github.com/Agarwalsahil/TodoAPI/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
