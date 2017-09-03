package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/TheoRev/OdontoSoft_Backend/controllers"
)

// SetUserRouter establece la ruta del modelo usuario
func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CreateUser).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
