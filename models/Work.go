package models

import (
	"github.com/jinzhu/gorm"
)

// Work estructura de la tabla trabajo
type Work struct {
	gorm.Model
	Name  string  `json:"name" gorm:"not null; type:varchar(150)"`
	Price float64 `json:"price" gorm:"not null; type:decimal(36,2)"`
}

// Works slice de trabajos
type Works []Work
