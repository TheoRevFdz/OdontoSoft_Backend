package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TheoRev/OdontoSoft_Backend/config"
	"github.com/TheoRev/OdontoSoft_Backend/models"
	"github.com/TheoRev/OdontoSoft_Backend/util"
)

// CreatePatient crea un nuevo paciente en la db
func CreatePatient(w http.ResponseWriter, r *http.Request) {
	patient := models.Patient{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el patient a registrar: %s", err)
		m.Code = http.StatusBadRequest
		util.DisplayMessage(w, m)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Create(&patient).Error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		m.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		m.Code = http.StatusBadRequest
		util.DisplayMessage(w, m)
		return
	}

	m.Message = "Paciente creado exitosamente"
	m.Code = http.StatusCreated
	// w.WriteHeader(http.StatusCreated)
	util.DisplayMessage(w, m)
}

// UpdatePatient actualiza un registro de paciente en la db
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	patient := models.Patient{}
	msg := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el paciente a actualizar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Save(&patient).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al actualizar el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Paciente actualizado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// DeletePatient elimina un paciente de la db
func DeletePatient(w http.ResponseWriter, r *http.Request) {
	patient := models.Patient{}
	msg := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el paciente a eliminar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	patient.State = false
	err = db.Save(&patient).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al eliminar el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Paciente eliminado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// FindAllPatients obtiene todos los pacientes de la db
func FindAllPatients(w http.ResponseWriter, r *http.Request) {
	patients := models.Patients{}
	msg := models.Message{}

	db := config.GetConnection()
	defer db.Close()

	// err := db.Where("state = TRUE").Order("id desc").Find(&patients).Error
	err := db.Order("id desc").Find(&patients).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al obtener los datos: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	j, err := json.Marshal(patients)
	if err != nil {
		log.Fatalf("Error al convertir los datos a json: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// FindPatientsWithoutTreatment obtiene todos los pacientes que no han iniciado tratamiento
func FindPatientsWithoutTreatment(w http.ResponseWriter, r *http.Request) {
	patients := models.Patients{}
	msg := models.Message{}

	db := config.GetConnection()
	defer db.Close()

	// err := db.Order("id desc").Find(&patients).Error
	err := db.Joins("left join treatments on treatments.patient_id = patients.id").Where("treatments.patient_id is null").Find(&patients).Error
	// err := db.Joins("left join treatments on treatments.id = patients.id").Find(&patients).Error

	if err != nil {
		msg.Message = fmt.Sprintf("Error al obtener los datos: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	j, err := json.Marshal(patients)
	if err != nil {
		log.Fatalf("Error al convertir los datos a json: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
