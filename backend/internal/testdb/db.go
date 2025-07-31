// backend/internal/testdb/db.go
package testdb

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Init retorna una DB SQLite en memoria ya conectada
func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	return db, err
}
