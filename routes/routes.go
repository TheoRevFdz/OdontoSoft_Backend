package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes inicializa todas las rutas de los modelos
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetUserRouter(router)
	SetLoginRouter(router)

	return router
}
