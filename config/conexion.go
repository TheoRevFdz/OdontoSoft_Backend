package config

import (
	"fmt"

	"github.com/TheoRev/OdontoSoft_Backend/models"
	"github.com/jinzhu/gorm"

	// Driver de postgresql
	_ "github.com/lib/pq"
)

var conexion *gorm.DB

// Connect establece la conexion con la db
func Connect() {
	url := GetUrlDatabase()
	connection, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	} else {
		conexion = connection
		fmt.Println("Conexión exitosa!")
	}
}

// CloseConnection cierra la conexion con la db
func CloseConnection() {
	conexion.Close()
}

// CreateTables crea tablas segun estructuras
func CreateTables() {
	conexion.DropTableIfExists(&models.Patient{})
	conexion.DropTableIfExists(&models.Treatment{})
	conexion.CreateTable(&models.Patient{})
	conexion.CreateTable(&models.Treatment{})
}
