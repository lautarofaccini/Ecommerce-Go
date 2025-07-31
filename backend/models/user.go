// backend/models/user.go
package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dataSourceName string) error {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return db.AutoMigrate(&User{})
}

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
