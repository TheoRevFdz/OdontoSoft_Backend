package main

import (
	"flag"
	"log"

	"github.com/TheoRev/OdontoSoft_Backend/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la DB")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Inició la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración.")
	}

	// db := config.GetConnection()
	// defer db.Close()

	// db.DropTableIfExists(&models.Patient{})
	// db.DropTableIfExists(&models.Treatment{})
	// db.CreateTable(&models.Patient{})
	// db.CreateTable(&models.Treatment{})
}
