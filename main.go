package main

import (
	"TaskApi/db"
	"TaskApi/models"
	"TaskApi/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)




func main() {
	//Conexión base de datos
	db.DBConnection() 
	//Creando tabla de Tareas
	db.DB.AutoMigrate(models.Task{})
	route:= mux.NewRouter()
	//Permitiendo CORS
	route.Use(mux.CORSMethodMiddleware(route))
	//Rutas
	route.HandleFunc("/tareas",routes.GetTasks).Methods("GET")
	route.HandleFunc("/tarea/{id}",routes.GetTaskInfo).Methods("GET") 
	route.HandleFunc("/crearTarea",routes.CreateTask).Methods("POST")
	route.HandleFunc("/update/{id}",routes.UpdateTask).Methods("PUT")
	route.HandleFunc("/delete/{id}",routes.DeleteTask).Methods("DELETE")

	//Levantando el Servidor
	fmt.Println("Escuchando en el puerto 3000")
	http.ListenAndServe("localhost:3000",route) 
}