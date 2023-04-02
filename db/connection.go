package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn="host=localhost user=postgres password=cryptocurrency dbname=gorm port=5432"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err =gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}else{
		log.Println("Base de datos conectada de forma correcta")
	}
}