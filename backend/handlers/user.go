// backend/handlers/user.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarofaccini/ecommerce-go/models"
)

func GetUsers(context *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	context.JSON(http.StatusOK, users)
}

func CreateUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&user)
	context.JSON(http.StatusCreated, user)
}

// TODO: EditUser, DeleteUser, GetUser
