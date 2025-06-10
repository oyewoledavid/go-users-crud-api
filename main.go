package main

import (
	"log"
	"users-api-gin/db"
	"users-api-gin/handlers"
	"users-api-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	db.DB.AutoMigrate(&models.User{}, &models.Account{})
	router := gin.Default()
	router.GET("/users", handlers.GetUser)
	router.GET("/users/id/:id", handlers.GetUserByID)
	router.GET("/users/name/:fullname", handlers.GetUserByName)
	router.POST("/users", handlers.AddUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	router.Run("localhost:8080")
}
