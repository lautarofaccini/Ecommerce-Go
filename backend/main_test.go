// backend/main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lautarofaccini/ecommerce-go/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter_RoutesExist(test *testing.T) {
	// Para no ver logs de Gin en tests
	gin.SetMode(gin.TestMode)

	// Inicializar DB en memoria
	if err := testhelpers.ConnectTest(); err != nil {
		test.Fatalf("No se pudo inicializar DB de test: %v", err)
	}

	router := SetupRouter()

	// Probamos que GET /users devuelva 200 (aunque sin DB conectada,
	// debería caer en Empty y devolver [] con OK)
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(test, http.StatusOK, recorder.Code, "GET /users debería existir y devolver 200")

	// Probamos que POST /users devuelva 400 (sin body válido)
	recorder2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodPost, "/users", nil)
	router.ServeHTTP(recorder2, req2)

	assert.Equal(test, http.StatusBadRequest, recorder2.Code,
		"POST /users sin JSON válido debería devolver 400 Bad Request")
}
