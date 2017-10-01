package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes inicializa todas las rutas de los modelos
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetUserRouter(router)
	SetLoginRouter(router)
	SetFindTreatmentsDetailByTreatmentIDRouter(router)
	SetCrudTreatmentDetailRouter(router)
	SetFindAllTreatmentsDetailRouter(router)
	SetCrudTreatmentRouter(router)
	SetFindPatientWhitoutTreatmentRouter(router)
	SetCrudPatientRouter(router)
	SetFindAllPatientRouter(router)
	SetFindAllTreatmentsRouter(router)
	SetLastTreatmentRouter(router)
	SetCrudControlRouter(router)
	SetFindAllControlsRouter(router)
	SetCrudWorkRouter(router)
	SetFindAllWorksRouter(router)
	SetLibsRouter(router)

	return router
}
