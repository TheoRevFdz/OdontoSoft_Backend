package models

import "github.com/jinzhu/gorm"

// Patient estructura de la tabla paciente
type Patient struct {
	gorm.Model
	DateInit     string `json:"dateInit" gorm:"not null; type:date"`
	NomApe       string `json:"nomApe" gorm:"not null; type:varchar(200)"`
	Age          uint   `json:"age" gorm:"not null; type:integer"`
	Sex          string `json:"sex" gorm:"not null; type:varchar(1)"`
	DateNac      string `json:"dateNac" gorm:"not null; type:date"`
	Address      string `json:"address" gorm:"type:varchar(200)"`
	Ocupation    string `json:"ocupation" gorm:"not null; type:varchar(100)"`
	TelCel       string `json:"telCel" gorm:"type:varchar(12)"`
	Alergies     string `json:"alergies" gorm:"type:varchar(150)"`
	Operations   string `json:"operations" gorm:"type:varchar(150)"`
	Diabettes    bool   `json:"diabettes" gorm:"not null; type:boolean"`
	Hipertension bool   `json:"hipertension" gorm:"not null; type:boolean"`
	Others       string `json:"others" gorm:"type:varchar(120)"`
	TreatMedics  string `json:"treatMedics" gorm:"type:varchar(120)"`
	State        bool   `json:"state" gorm:"not null; type:boolean; default:TRUE"`
}

// Patients slice de pacientes
type Patients []Patient
