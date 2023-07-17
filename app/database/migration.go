package database

import (
	"fmt"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(

	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}