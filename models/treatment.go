package models

import "github.com/jinzhu/gorm"

// Treatment estructura de la tabla tratamiento
type Treatment struct {
	gorm.Model
	Work      string  `json:"work" gorm:"type: varchar(100)"`
	Quantity  uint    `json:"quantity" gorm:"not null; type:integer"`
	Cost      float64 `gorm:"not null; type:decimal(36,2)"`
	IdPatient uint    `json:"idPatient" gorm:"not null"`
}

// Treatments slice de tratamientos
type Treatments []Treatment
