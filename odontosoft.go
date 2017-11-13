package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/urfave/negroni"

	"github.com/TheoRev/OdontoSoft_Backend/migration"
	"github.com/TheoRev/OdontoSoft_Backend/routes"
	"github.com/TheoRev/OdontoSoft_Backend/util"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la DB")
	flag.IntVar(&util.Port, "port", 3030, "Puerto para el servidor web")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Inició la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración.")
	}

	router := routes.InitRoutes()
	neg := negroni.Classic()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
	neg.Use(c)
	neg.UseHandler(router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", util.Port),
		Handler: neg,
	}

	log.Printf("Iniciado el servidor en http://localhost:%d", util.Port)
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución del programa")
}
