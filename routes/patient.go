package routes

import (
	"github.com/TheoRev/OdontoSoft_Backend/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCreatePatientRouter estaablece la ruta para crear pacientes
func SetCreatePatientRouter(router *mux.Router) {
	prefix := "/api/create/patients"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CreatePatient).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetFindAllPatientRouter estaablece la ruta para obtener todos los pacientes
func SetFindAllPatientRouter(router *mux.Router) {
	prefix := "/api/patients"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.FindAllPatients).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
