package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TheoRev/OdontoSoft_Backend/config"
	"github.com/TheoRev/OdontoSoft_Backend/models"
	"github.com/TheoRev/OdontoSoft_Backend/util"
	"github.com/gorilla/mux"
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

// FindTreatmentsDetailByTreatmentID obtiene todos los tratamientos de la db
func FindTreatmentsDetailByTreatmentID(w http.ResponseWriter, r *http.Request) {
	tsd := models.TreatmentsDetail{}
	msg := models.Message{}

	vars := mux.Vars(r)
	treatmentID, _ := strconv.Atoi(vars["id"])

	db := config.GetConnection()
	defer db.Close()

	err := db.Where("treatment_id = ?", treatmentID).Find(&tsd).Error
	for i := 0; i < len(tsd); i++ {
		db.Model(&tsd[i]).Related(&tsd[i].Work)
	}

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

// FindAllTreatmentsDetail obtiene todos los tratamientos de la db
func FindAllTreatmentsDetail(w http.ResponseWriter, r *http.Request) {
	tsd := models.TreatmentsDetail{}
	msg := models.Message{}

	db := config.GetConnection()
	defer db.Close()

	err := db.Find(&tsd).Error
	for i := 0; i < len(tsd); i++ {
		db.Model(&tsd[i]).Related(&tsd[i].Work)
	}

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
