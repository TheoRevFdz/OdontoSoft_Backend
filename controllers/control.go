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

// CreateControl crea un nuevo tratamiento en la db
func CreateControl(w http.ResponseWriter, r *http.Request) {
	control := models.Control{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&control)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el control a registrar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Create(&control).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Control creado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// UpdateControl crea un nuevo tratamiento en la db
func UpdateControl(w http.ResponseWriter, r *http.Request) {
	control := models.Control{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&control)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el control a actualizar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Save(&control).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al actualizar el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Control actualizado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// DeleteControl crea un nuevo tratamiento en la db
func DeleteControl(w http.ResponseWriter, r *http.Request) {
	control := models.Control{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&control)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el control a eliminar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Delete(&control).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al eliminar el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Control eliminado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// FindAllControls obtiene todos los tratamientos de la db
func FindAllControls(w http.ResponseWriter, r *http.Request) {
	controls := models.Controls{}
	msg := models.Message{}

	db := config.GetConnection()
	defer db.Close()

	err := db.Find(&controls).Error
	// err := db.Model(&controls).Association("Treatment").Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al obtener los datos: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	j, err := json.Marshal(controls)
	if err != nil {
		log.Fatalf("Error al convertir los datos a json: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
