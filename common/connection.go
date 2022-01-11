package common

import (
	"log"

	"github.com/Camiloxrc/Point-of-sale/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Obtener la conexion a nuestra base de datos
func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:0ee74af421@/Alondra?charset=utf8&parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

//Funcion para migrar la base de datos
func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Escribiendo datos")

	db.AutoMigrate(&models.Cliente{})
}
