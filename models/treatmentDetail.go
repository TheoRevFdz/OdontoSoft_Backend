package models

// TreatmentDetail estructura q contienes todas las curaciones realizadas al paciente
type TreatmentDetail struct {
	WorkID      uint   `json:"workId" gorm:"type: integer"`
	Works       []Work `json:"works" gorm:"ForeignKey:TreatmentDetailID;AssociationForeignKey:WorkID"`
	Quantity    uint   `json:"quantity" gorm:"not null; type:integer"`
	TreatmentID uint   `json:"treatmentId"`
}
