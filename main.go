package main

import (
	"app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/fruits", handlers.GetFruits)
	router.GET("/fruits/:id", handlers.GetFruitByID)
	router.POST("/fruits/create", handlers.CreateFruit)

	router.Run("localhost:8080") // listen and serve on localhost:8080
}
