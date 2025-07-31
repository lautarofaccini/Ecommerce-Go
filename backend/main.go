// backend/main.go
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lautarofaccini/ecommerce-go/handlers"
	"github.com/lautarofaccini/ecommerce-go/models"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	usersGroup := router.Group("/users")
	{
		usersGroup.GET("", handlers.GetUsers)
		usersGroup.POST("", handlers.CreateUser)
		// PUT, DELETE, GET("/:id")…
	}
	return router
}

func main() {
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL no está definida")
	}
	if err := models.Connect(dbURL); err != nil {
		panic(err)
	}

	router := SetupRouter()
	router.Run(":8080")
}
