package database

import (
	"fmt"
	author "online-library/features/authors/data"
	book "online-library/features/books/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&book.Book{},
		&author.Author{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
