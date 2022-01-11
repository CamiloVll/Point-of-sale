package routes

import (
	"github.com/Camiloxrc/Point-of-sale/controller"
	"github.com/gorilla/mux"
)

// Rutas de la api
func SetClienteRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/cliente/api").Subrouter()
	subRoute.HandleFunc("/all", controller.GetAll).Methods("GET")
	subRoute.HandleFunc("/find/{id}", controller.Get).Methods("GET")
	subRoute.HandleFunc("/save", controller.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controller.Delete).Methods("POST")
}
