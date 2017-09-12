package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetLibsRouter estaablece la ruta para crear pacientes
func SetLibsRouter(router *mux.Router) {
	staticFiles := http.FileServer(http.Dir("assets"))
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	// subRouter.HandleFunc("/", controllers.CreatePatient).Methods("POST")
	subRouter.Handler

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
