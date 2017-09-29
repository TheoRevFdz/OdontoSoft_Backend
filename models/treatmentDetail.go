package models

import (
	"github.com/jinzhu/gorm"
)

// TreatmentDetail estructura q contienes todas las curaciones realizadas al paciente
type TreatmentDetail struct {
	gorm.Model
	WorkID      uint `json:"workId" gorm:"type: integer"`
	Work        Work `json:"works" gorm:"ForeignKey:TreatmentDetailID;AssociationForeignKey:WorkID"`
	Quantity    uint `json:"quantity" gorm:"not null; type:integer"`
	TreatmentID uint `json:"treatmentId" gorm:"type: integer"`
}

// TreatmentsDetail slice de curaciones de los tratamientos
type TreatmentsDetail []TreatmentDetail
