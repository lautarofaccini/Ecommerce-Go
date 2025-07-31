// backend/models/user_test.go
// unit test
package models

import (
	"testing"

	"github.com/lautarofaccini/ecommerce-go/internal/testdb"
	"github.com/stretchr/testify/assert"
)

func TestUserModel_CreateAndFetch(test *testing.T) {
	// Conectar DB in-memory
	db, err := testdb.Init()
	assert.NoError(test, err)
	DB = db
	err = db.AutoMigrate(&User{})
	assert.NoError(test, err)

	// Crear un usuario
	user := User{Name: "Test", Email: "test@example.com"}
	result := DB.Create(&user)
	assert.NoError(test, result.Error)
	assert.NotZero(test, user.ID)

	// Leer el usuario
	var user2 User
	err = DB.First(&user2, user.ID).Error
	assert.NoError(test, err)
	assert.Equal(test, user.Name, user2.Name)
}
