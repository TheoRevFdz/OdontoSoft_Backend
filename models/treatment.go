package models

import "github.com/jinzhu/gorm"

// Treatment estructura de la tabla tratamiento
type Treatment struct {
	gorm.Model
	IDWork    uint `json:"idWork" gorm:"type: integer"`
	Quantity  uint `json:"quantity" gorm:"not null; type:integer"`
	IDPatient uint `json:"idPatient" gorm:"not null"`
}

// Treatments slice de tratamientos
type Treatments []Treatment
