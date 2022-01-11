package main

import (
	"log"
	"net/http"

	"github.com/Camiloxrc/Point-of-sale/common"
	"github.com/Camiloxrc/Point-of-sale/routes"
	"github.com/gorilla/mux"
)

func main() {
	common.Migrate()

	router := mux.NewRouter()
	routes.SetClienteRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}
	log.Println("Servidor corrriendo en el puerto 9000")
	log.Println(server.ListenAndServe())
}
