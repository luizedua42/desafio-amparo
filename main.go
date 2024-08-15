package main

import (
	gin "github.com/gin-gonic/gin"
	"desafio-amparo/handlers"
	// "desafio-amparo/models"
)

func main() {
	router := gin.Default()

	router.GET("/tasks", handlers.GetTasks)
	router.DELETE("/tasks/:id", handlers.DelTask)
	router.PUT("/tasks", handlers.UpdateTask)
	router.Run(":8080")
}