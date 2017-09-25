package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes inicializa todas las rutas de los modelos
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// SetLoginGUI(router)
	SetUserRouter(router)
	SetLoginRouter(router)
	SetCrudPatientRouter(router)
	SetFindAllPatientRouter(router)
	SetCrudTreatmentRouter(router)
	SetFindAllTreatmentsRouter(router)
	SetCrudControlRouter(router)
	SetFindAllControlsRouter(router)
	SetCrudWorkRouter(router)
	SetFindAllWorksRouter(router)
	SetLibsRouter(router)

	return router
}
