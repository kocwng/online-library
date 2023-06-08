package data

import (

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name    string
	Country string
}
