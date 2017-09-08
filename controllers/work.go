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

// CreateWork crea un nuevo trabajo en la db
func CreateWork(w http.ResponseWriter, r *http.Request) {
	work := models.Work{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&work)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el trabajo a registrar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Create(&work).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Trabajo creado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// UpdateWork actualiza un trabajo en la db
func UpdateWork(w http.ResponseWriter, r *http.Request) {
	work := models.Work{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&work)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el trabajo a actualizar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Save(&work).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al actualizar el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Trabajo actualizado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// DeleteWork elimina un trabajo en la db
func DeleteWork(w http.ResponseWriter, r *http.Request) {
	work := models.Work{}
	msg := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&work)
	if err != nil {
		msg.Message = fmt.Sprintf("Error al leer el trabajo a eliminar: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	err = db.Delete(&work).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al eliminar el registro: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	msg.Message = "Trabajo eliminado con éxito"
	msg.Code = http.StatusCreated
	util.DisplayMessage(w, msg)
}

// FindAllWork obtiene todos los trabajos de la db
func FindAllWork(w http.ResponseWriter, r *http.Request) {
	works := models.Works{}
	msg := models.Message{}

	db := config.GetConnection()
	defer db.Close()

	err := db.Find(&works).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al obtener los datos: %s", err)
		msg.Code = http.StatusBadRequest
		util.DisplayMessage(w, msg)
		return
	}

	j, err := json.Marshal(works)
	if err != nil {
		log.Fatalf("Error al convertir los datos a json: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
