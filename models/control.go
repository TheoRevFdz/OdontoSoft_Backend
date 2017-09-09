package models

import (
	"github.com/jinzhu/gorm"
)

// Control estructura de la tabla control
type Control struct {
	gorm.Model
	Fecha       string  `json:"fecha" gorm:"not null; type:date; default: NOW()"`
	IDWork      uint    `json:"idWork" gorm:"not null"`
	Pieces      uint    `json:"pieces" gorm:"not null; type:integer; default:0"`
	Description string  `json:"description" gorm:"type:varchar(150)"`
	Pago        float64 `json:"pago" gorm:"type:decimal(36,2); default:0.00"`
	IDTreatment uint    `json:"idTreatment" gorm:"not null"`
	Total       float64 `json:"total" gorm:"-"`
}

// Controls slice de control
type Controls []Control
