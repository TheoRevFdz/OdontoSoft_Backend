package models

import "github.com/jinzhu/gorm"

// Treatment estructura de la tabla tratamiento
type Treatment struct {
	gorm.Model
	PatientID         uint              `json:"patientId" gorm:"not null;type: integer"`
	Patient           Patient           `json:"patient" gorm:"ForeignKey:TreatmentID;AssociationForeignKey:PatientID"`
	TreatmentDetailID uint              `json:"treatmentDetailId"`
	TreatmentsDetail  []TreatmentDetail `json:"treatmenetsDetail" gorm:"ForeignKey:TreatmentID;AssociationForeignKey:TreatmentDetailID"`
	Select            bool              `json:"select" gorm:"-"`
}

// Treatments slice de tratamientos
type Treatments []Treatment
