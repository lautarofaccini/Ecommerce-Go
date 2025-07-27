// /backend/main.go
package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/lautarofaccini/ecommerce-go/handlers"
	"github.com/lautarofaccini/ecommerce-go/models"
)

func main (){
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL no está definida")
	}
	if err := models.Connect(dsn); err != nil {
		panic(err)
	}

	r := gin.Default()
	users := r.Group("/users")
	{
		users.GET("", handlers.GetUsers)
		users.POST("", handlers.CreateUser)
		// PUT, DELETE, GET
	}
	r.Run(":8080")
}
