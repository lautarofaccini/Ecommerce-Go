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

	dataSourceName := os.Getenv("DATABASE_URL")
	if dataSourceName == "" {
		log.Fatal("DATABASE_URL no está definida")
	}
	if err := models.Connect(dataSourceName); err != nil {
		panic(err)
	}

	router := SetupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
