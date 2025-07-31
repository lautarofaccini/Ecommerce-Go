// backend/handlers/user_test.go
// integration tests
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lautarofaccini/ecommerce-go/models"
	"github.com/lautarofaccini/ecommerce-go/testhelpers"
	"github.com/stretchr/testify/assert"
)

// setupRouter prepara Gin con routes y base en memoria
func setupRouter() *gin.Engine {
	// inicializa DB test
	if err := testhelpers.ConnectTest(); err != nil {
		panic(err)
	}

	router := gin.Default()
	users := router.Group("/users")
	{
		users.GET("", GetUsers)
		users.POST("", CreateUser)
	}
	return router
}

func TestGetUsers_Empty(test *testing.T) {
	test.Parallel()
	router := setupRouter()

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(test, http.StatusOK, recorder.Code)
	var res []models.User
	err := json.Unmarshal(recorder.Body.Bytes(), &res)
	assert.NoError(test, err)
	assert.Empty(test, res, "Debe retornar lista vacía al inicio")
}

func TestCreateUser_And_Get(test *testing.T) {
	router := setupRouter()

	// Crear usuario
	payload := map[string]string{"name": "Test", "email": "test@example.com"}
	fmt.Println(payload)
	body, _ := json.Marshal(payload)

	// Crear variables de la petición
	recorder1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/users", bytes.NewReader(body))
	req1.Header.Set("Content-Type", "application/json")

	// Simular petición y extraer usuario
	router.ServeHTTP(recorder1, req1)
	var user models.User
	err := json.Unmarshal(recorder1.Body.Bytes(), &user)

	// Tests
	assert.Equal(test, http.StatusCreated, recorder1.Code)
	assert.NoError(test, err)
	assert.Equal(test, uint(1), user.ID)
	assert.Equal(test, "Test", user.Name)
	assert.Equal(test, "test@example.com", user.Email)

	// Obtener lista y verificar que incluye al creado
	recorder2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/users", nil)

	// Simular petición
	router.ServeHTTP(recorder2, req2)
	var list []models.User
	err = json.Unmarshal(recorder2.Body.Bytes(), &list)

	// Tests
	assert.Equal(test, http.StatusOK, recorder2.Code)
	assert.NoError(test, err)
	assert.Len(test, list, 1)
	assert.Equal(test, user, list[0])
}
