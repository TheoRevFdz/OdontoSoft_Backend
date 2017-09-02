package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Control estructura de la tabla control
type Control struct {
	gorm.Model
	Fecha       time.Time `json:"fecha" gorm:"not null; type:date"`
	IDWork      uint      `json:"idWork" gorm:"not null"`
	Piece       string    `json:"piece" gorm:"not null; type:varchar(4)"`
	ACuenta     float64   `json:"aCuenta" gorm:"type:decimal(36,2); default:0.00"`
	IDTreatment uint      `json:"idTreatment" gorm:"not null"`
}

// Controls slice de control
type Controls []Control
