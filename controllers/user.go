package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TheoRev/OdontoSoft_Backend/config"
	"github.com/TheoRev/OdontoSoft_Backend/models"
	"github.com/TheoRev/OdontoSoft_Backend/util"
)

// CreateUser registra un nuevo usuario en la db
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

	db := config.GetConnection()
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

// Login autentifica el acceso de usuario al sistema
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	cod := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", cod)

	db.Where("username=? and password=?", user.Username, pwd).First(&user)

	if user.ID > 0 {
		user.Password = ""
		j, err := json.Marshal(user)
		if err != nil {
			log.Fatalf("Error al convertir el token a json: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message: "Usuario o clave no válido",
			Code:    http.StatusUnauthorized,
		}
		util.DisplayMessage(w, m)
	}
}
