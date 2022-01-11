package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Camiloxrc/Point-of-sale/common"
	"github.com/Camiloxrc/Point-of-sale/models"
	"github.com/gorilla/mux"
)

// Funcion para optener varios registros
func GetAll(writer http.ResponseWriter, request *http.Request) {
	clientes := []models.Cliente{}
	db := common.GetConnection()
	defer db.Close()

	db.Find(&clientes)

	json, _ := json.Marshal(clientes)
	common.SendResponse(writer, http.StatusOK, json)
}

// Funcion para obtener un registro
func Get(writer http.ResponseWriter, request *http.Request) {
	cliente := models.Cliente{}

	id := mux.Vars(request)["id"]

	db := common.GetConnection()
	defer db.Close()

	db.Find(&cliente, id)

	if cliente.ID > 0 {
		json, _ := json.Marshal(cliente)
		common.SendResponse(writer, http.StatusOK, json)
	} else {
		common.SendError(writer, http.StatusNotFound)
	}
}

// Funcion para guardar los cambios
func Save(writer http.ResponseWriter, request *http.Request) {
	cliente := models.Cliente{}
	db := common.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&cliente)

	if error != nil {
		log.Fatal(error)
		common.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&cliente).Error

	if error != nil {
		log.Fatal(error)
		common.SendError(writer, http.StatusInternalServerError)
		return
	}
	json, _ := json.Marshal(cliente)

	common.SendResponse(writer, http.StatusCreated, json)
}

// Funcion para borrar registros
func Delete(writer http.ResponseWriter, request *http.Request) {
	cliente := models.Cliente{}
	db := common.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&cliente, id)

	if cliente.ID > 0 {
		db.Delete(cliente)
		common.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(writer, http.StatusNotFound)
	}
}
