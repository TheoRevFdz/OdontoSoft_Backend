package models

import "github.com/jinzhu/gorm"

// Treatment estructura de la tabla tratamiento
type Treatment struct {
	gorm.Model
	WorkID    uint    `json:"workId" gorm:"type: integer"`
	Works     []Work  `json:"works" gorm:"ForeignKey:TreatmentID;AssociationForeignKey:WorkID"`
	Quantity  uint    `json:"quantity" gorm:"not null; type:integer"`
	PatientID uint    `json:"patientId" gorm:"not null;type: integer"`
	Patient   Patient `json:"patient" gorm:"ForeignKey:TreatmentID;AssociationForeignKey:PatientID"`
}

// Treatments slice de tratamientos
type Treatments []Treatment
