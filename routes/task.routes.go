package routes

import (
	"TaskApi/db"
	"TaskApi/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//getTasks - GET - /getTasks
func GetTasks(res http.ResponseWriter, req *http.Request){
	var tasks []models.Task
	db.DB.Find(&tasks)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(tasks)
	
}

func GetTaskInfo(res http.ResponseWriter, req *http.Request){
	var task models.Task
	params:=mux.Vars(req)
	db.DB.First(&task,params["id"])
	if task.ID == 0{
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Tarea no encontrada"))
		return
	}
	json.NewEncoder(res).Encode(&task)
}

func CreateTask(res http.ResponseWriter, req *http.Request){
	var task models.Task
	json.NewDecoder(req.Body).Decode(&task)

	createdTask:=db.DB.Create(&task)
	err:= createdTask.Error
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
	}
	json.NewEncoder(res).Encode(&task)
}

func UpdateTask(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res,"<h1>Hola Mundo desde la ruta PUT</h1>")
}

func DeleteTask(res http.ResponseWriter, req *http.Request){
	var task models.Task
	params:=mux.Vars(req)
	db.DB.First(&task,params["id"])
	if task.ID == 0 {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Tarea no encontrada"))
		return
	}
		db.DB.Unscoped().Delete(&task)
		res.WriteHeader(http.StatusAccepted)
		res.Write([]byte("Tarea eleminada correctamente"))
}