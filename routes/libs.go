package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetLibsRouter estaablece la ruta para crear pacientes
func SetLibsRouter(router *mux.Router) {
	prefix := "/"
	staticFiles := http.FileServer(http.Dir("assets"))
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(false)
	// subRouter.HandleFunc("/", controllers.CreatePatient).Methods("POST")
	subRouter.Handle("/assets/", http.StripPrefix("/assets/", staticFiles))

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
