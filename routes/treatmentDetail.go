package routes

import (
	"github.com/TheoRev/OdontoSoft_Backend/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCrudTreatmentDetailRouter estaablece la ruta para crear un tratamiento
func SetCrudTreatmentDetailRouter(router *mux.Router) {
	prefix := "/api/crud/treatment-detail"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CreateTreatmentDetail).Methods("POST")
	subRouter.HandleFunc("/", controllers.UpdateTreatmentDetail).Methods("PUT")
	subRouter.HandleFunc("/", controllers.DeleteTreatmentDetail).Methods("DELETE")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetFindAllTreatmentsDetailRouter estaablece la ruta para obtener todos las curaciones de los tratamientos
func SetFindAllTreatmentsDetailRouter(router *mux.Router) {
	prefix := "/api/treatments-detail"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.FindAllTreatmentsDetail).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetFindTreatmentsDetailByTreatmentIDRouter estaablece la ruta para obtener todos las curaciones de los tratamientos
func SetFindTreatmentsDetailByTreatmentIDRouter(router *mux.Router) {
	prefix := "/api/treatments-detail/{id:[0-9]+}"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.FindTreatmentsDetailByTreatmentID).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
