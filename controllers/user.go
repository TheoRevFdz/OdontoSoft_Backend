package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheoRev/gocomments/configuration"

	"github.com/TheoRev/OdontoSoft_Backend/models"
	"github.com/TheoRev/OdontoSoft_Backend/util"
)

// CreateUseer registra un nuevo usuario en la db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario a registrar: %s", err)
		m.Code = http.StatusBadRequest
		util.DisplayMessage(w, m)
		return
	}

	cod := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", cod)
	user.Password = pwd

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&user).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		m.Code = http.StatusBadRequest
		util.DisplayMessage(w, m)
		return
	}

	m.Message = "Usuario creado con éxito"
	m.Code = http.StatusCreated
	util.DisplayMessage(w, m)
}
