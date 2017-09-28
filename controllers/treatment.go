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

// CreateTreatment crea un nuevo tratamiento en la db
func CreateTreatment(w http.ResponseWriter, r *http.Request) {
	treatment := models.Treatment{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&treatment)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el tratamiento a registrar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Create(&treatment).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Tratamiento creado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// UpdateTreatment actualiza un registro de tratamiento en la db
func UpdateTreatment(w http.ResponseWriter, r *http.Request) {
	treatment := models.Treatment{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&treatment)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el tratamiento a actualizar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Save(&treatment).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al actualizar el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Tratamiento actualizado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// DeleteTreatment elimina un registro de tratamiento en la db
func DeleteTreatment(w http.ResponseWriter, r *http.Request) {
	treatment := models.Treatment{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&treatment)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el tratamiento a eliminar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Delete(&treatment).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al eliminar el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Tratamiento eliminado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// FindAllTreatments obtiene todos los tratamientos de la db
func FindAllTreatments(w http.ResponseWriter, r *http.Request) {
	treatments := models.Treatments{}
	msg := models.Message{}

	db := config.GetConnection()
	defer db.Close()

	err := db.Find(&treatments).Error
	for i := 0; i < len(treatments); i++ {
		db.Model(&treatments[i]).Related(&treatments[i].Patient)
	}

	if err != nil {
		msg.Message = fmt.Sprintf("Error al obtener los datos: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	j, err := json.Marshal(treatments)
	if err != nil {
		log.Fatalf("Error al convertir los datos a json: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// FindLastTreatment obtiene el ultimo tratamiento registrado en la db
func FindLastTreatment(w http.ResponseWriter, r *http.Request) {
	treatment := models.Treatment{}
	msg := models.Message{}

	db := config.GetConnection()
	defer db.Close()

	err := db.Last(&treatment).Error

	if err != nil {
		msg.Message = fmt.Sprintf("Error al obtener los datos: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	j, err := json.Marshal(treatment)
	if err != nil {
		log.Fatalf("Error al convertir los datos a json: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
