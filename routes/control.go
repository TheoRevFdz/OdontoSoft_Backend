package routes

import (
	"github.com/TheoRev/OdontoSoft_Backend/controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCrudControlRouter estaablece la ruta para crear control del tratamiento
func SetCrudControlRouter(router *mux.Router) {
	prefix := "/api/crud/control"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CreateControl).Methods("POST")
	subRouter.HandleFunc("/", controllers.UpdateControl).Methods("PUT")
	subRouter.HandleFunc("/", controllers.DeleteControl).Methods("DELETE")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetFindAllControlsRouter estaablece la ruta para obtener todos los controles de tratamiento
func SetFindAllControlsRouter(router *mux.Router) {
	prefix := "/api/controls"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.FindAllControls).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
