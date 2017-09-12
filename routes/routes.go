package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes inicializa todas las rutas de los modelos
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetUserRouter(router)
	SetLoginGUI(router)
	SetLoginRouter(router)
	SetCrudPatientRouter(router)
	SetFindAllPatientRouter(router)
	SetCrudTreatmentRouter(router)
	SetFindAllTreatmentsRouter(router)
	SetCrudControlRouter(router)
	SetFindAllControlsRouter(router)
	SetCrudWorkRouter(router)
	SetFindAllWorksRouter(router)

	return router
}
