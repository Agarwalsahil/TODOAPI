package main

import (
	"log"
	"os"

	"github.com/Agarwalsahil/TodoAPI/controllers"
	"github.com/Agarwalsahil/TodoAPI/db"
	"github.com/Agarwalsahil/TodoAPI/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	db.InitDB()

	server := gin.Default()

	server.POST("/register", controllers.Register)
	server.POST("/login", controllers.Login)

	authenticated := server.Group("/")
	authenticated.Use(middleware.AuthMiddleware)

	authenticated.POST("/todos", controllers.CreateTodo)
	authenticated.GET("/todos", controllers.GetTodos)
	authenticated.PUT("/todos/:id", controllers.UpdateTodo)
	authenticated.DELETE("/todos/:id", controllers.DeleteTodo)

	log.Println("Server running on http://localhost:8080")
	server.Run(":" + port)

}
