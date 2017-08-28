package models

import "github.com/jinzhu/gorm"

// Treatment estructura de la tabla tratamiento
type Treatment struct {
	gorm.Model
	TreatmentID uint    `json:"treatmentId" gorm:"type: serial; primary_key"`
	Work        string  `json:"work" gorm:"type: varchar(100)"`
	Quantity    int     `json:"quantity"`
	Cost        float64 `gorm:"type:decimal(36,2)"`
	CodPatient  uint    `json:"codPatient"`
}

// Treatments slice de tratamientos
type Treatments []Treatment
