package routes

import (
	"github.com/TheoRev/OdontoSoft_Backend/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCrudWorkRouter estaablece la ruta para crear un trabajo
func SetCrudWorkRouter(router *mux.Router) {
	prefix := "/api/crud/work"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CreateWork).Methods("POST")
	subRouter.HandleFunc("/", controllers.UpdateWork).Methods("PUT")
	subRouter.HandleFunc("/", controllers.DeleteWork).Methods("DELETE")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetFindAllWorksRouter estaablece la ruta para obtener todos los trabajos
func SetFindAllWorksRouter(router *mux.Router) {
	prefix := "/api/works"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.FindAllWork).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
