package routes

import (
	"github.com/TheoRev/OdontoSoft_Backend/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCrudTreatmentRouter estaablece la ruta para crear un tratamiento
func SetCrudTreatmentRouter(router *mux.Router) {
	prefix := "/api/crud/treatment"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CreateTreatment).Methods("POST")
	subRouter.HandleFunc("/", controllers.UpdateTreatment).Methods("PUT")
	subRouter.HandleFunc("/", controllers.DeleteTreatment).Methods("DELETE")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetFindAllTreatmentsRouter estaablece la ruta para obtener todos los tratamientos
func SetFindAllTreatmentsRouter(router *mux.Router) {
	prefix := "/api/controls"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.FindAllTreatments).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
