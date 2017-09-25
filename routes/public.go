package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetLibsRouter expone los archivos estaticos
func SetLibsRouter(router *mux.Router) {
	router.Handle("/", http.FileServer(http.Dir("./public")))
}
