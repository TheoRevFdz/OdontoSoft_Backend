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

// CreateTreatmentDetail crea un registro de curaciones
func CreateTreatmentDetail(w http.ResponseWriter, r *http.Request) {
	td := models.TreatmentDetail{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&td)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer la curación a registrar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Create(&td).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Curación creada con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// UpdateTreatmentDetail crea un registro de curaciones
func UpdateTreatmentDetail(w http.ResponseWriter, r *http.Request) {
	td := models.TreatmentDetail{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&td)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer la curación a actualizar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Save(&td).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Curación actualizada con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// DeleteTreatmentDetail crea un registro de curaciones
func DeleteTreatmentDetail(w http.ResponseWriter, r *http.Request) {
	td := models.TreatmentDetail{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&td)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer la curación a eliminar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Delete(&td).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Curación eliminada con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// FindAllTreatments obtiene todos los tratamientos de la db
func FindAllTreatmentsDetail(w http.ResponseWriter, r *http.Request) {
	tsd := models.TreatmentsDetail{}
	msg := models.Message{}

	db := config.GetConnection()
	defer db.Close()

	err := db.Find(&tsd).Error

	if err != nil {
		msg.Message = fmt.Sprintf("Error al obtener los datos: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	j, err := json.Marshal(tsd)
	if err != nil {
		log.Fatalf("Error al convertir los datos a json: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
