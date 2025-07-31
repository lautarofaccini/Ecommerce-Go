// backend/testhelpers/test_db.go
package testhelpers

import (
	"github.com/lautarofaccini/ecommerce-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectTest inicializa DB en memoria para tests
func ConnectTest() error {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return err
	}
	models.DB = db
	return db.AutoMigrate(&models.User{})
}
