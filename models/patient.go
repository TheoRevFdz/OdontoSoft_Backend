package models

import "time"
import "github.com/jinzhu/gorm"

// Patient estructura de la tabla paciente
type Patient struct {
	gorm.Model
	PatientID              uint      `json:"patientId"`
	FechaIngreso    time.Time `json:"fechaIngreso"`
	NombreApellidos string
	Edad            int
	Sexo            string
	FechaNacimiento time.Time
	Domicilio       string
	Ocupacion       string
	TelCel          string
	Alergias        string
	Operaciones     string
	Diabettes       bool
	Hipertension    bool
	Otras           string
	TratMedicos     string
	Estado          bool
}

// Patients slice de pacientes
type Patients []Patient
