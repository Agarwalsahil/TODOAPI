package controllers

import (
	"net/http"

	"github.com/Agarwalsahil/TodoAPI/db"
	"github.com/Agarwalsahil/TodoAPI/models"
	"github.com/Agarwalsahil/TodoAPI/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User

	result := db.DB.Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password){
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Could not generate token"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"token": token})
}
