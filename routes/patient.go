package routes

import (
	"github.com/TheoRev/OdontoSoft_Backend/controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCrudPatientRouter estaablece la ruta para crear pacientes
func SetCrudPatientRouter(router *mux.Router) {
	prefix := "/api/crud/patients"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.Headers("Access-Control-Allow-Origin", "*")
	subRouter.HandleFunc("/", controllers.CreatePatient).Methods("POST")
	subRouter.HandleFunc("/", controllers.CreatePatient).Methods("OPTIONS")
	subRouter.HandleFunc("/", controllers.UpdatePatient).Methods("PUT")
	subRouter.HandleFunc("/", controllers.DeletePatient).Methods("DELETE")

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

// SetFindPatientWhitoutTreatmentRouter estaablece la ruta para obtener los pacientes que no han iniciado un tratamiento
func SetFindPatientWhitoutTreatmentRouter(router *mux.Router) {
	prefix := "/api/patients-whitout-treatment"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.FindPatientsWithoutTreatment).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
