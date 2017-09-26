package migration

import (
	"github.com/TheoRev/OdontoSoft_Backend/config"
	"github.com/TheoRev/OdontoSoft_Backend/models"
)

// Migrate crea las tablas del modelo de db
func Migrate() {
	db := config.GetConnection()
	defer db.Close()

	db.DropTableIfExists(&models.User{})
	db.DropTableIfExists(&models.Patient{})
	db.DropTableIfExists(&models.Treatment{})
	db.DropTableIfExists(&models.TreatmentDetail{})
	db.DropTableIfExists(&models.Control{})
	db.DropTableIfExists(&models.Work{})

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Patient{})
	db.CreateTable(&models.Treatment{})
	db.CreateTable(&models.TreatmentDetail{})
	db.CreateTable(&models.Control{})
	db.CreateTable(&models.Work{})
}
