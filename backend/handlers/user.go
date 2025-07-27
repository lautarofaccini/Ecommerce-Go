// backend/handlers/user.go
package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/lautarofaccini/ecommerce-go/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context){
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&u)
	c.JSON(http.StatusCreated, u)
}

// TODO: EditUser, DeleteUser, GetUser
