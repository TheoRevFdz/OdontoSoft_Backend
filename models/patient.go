package models

import "time"

// Patient estructura de la tabla paciente
type Patient struct {
	ID              int64
	FechaIngreso    time.Time
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
	Estado          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Patients slice de pacientes
type Patients []Patient
