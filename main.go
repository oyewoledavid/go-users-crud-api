package main

import (
	"github.com/joho/godotenv"
	"log"
	"users-api-gin/handlers"
	"github.com/gin-gonic/gin"
	"users-api-gin/db"	
)

func main(){
	err := godotenv.Load()
	if err != nil {
	log.Fatal("Error loading .env file")
	}

	db.InitDB()
	
	router := gin.Default()
	router.GET("/users", handlers.GetUser)
	router.GET("/users/id/:id", handlers.GetUserByID)
	router.GET("/users/name/:fullname", handlers.GetUserByName)
	router.POST("/users", handlers.AddUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	router.Run("localhost:8080")
}
