package migration

import (
	"github.com/TheoRev/OdontoSoft_Backend/config"
	"github.com/TheoRev/OdontoSoft_Backend/models"
)

// Migrate crea las tablas del modelo de db
func Migrate() {
	db := config.GetConnection()
	defer db.Close()

	db.CreateTable(&models.Patient{})
	db.CreateTable(&models.Treatment{})
}
