// backend/testhelpers/test_db.go
package testhelpers

import (
	"github.com/lautarofaccini/ecommerce-go/internal/testdb"
	"github.com/lautarofaccini/ecommerce-go/models"
)

// ConnectTest inicializa DB en memoria para tests
func ConnectTest() error {
	db, err := testdb.Init()
	if err != nil {
		return err
	}
	models.DB = db
	return db.AutoMigrate(&models.User{})
}
