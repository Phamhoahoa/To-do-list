package controller

import(
	"net/http"
	"TODOLIST/model"
	"TODOLIST/dao"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"fmt"
)

func InsertTask(w http.ResponseWriter, r *http.Request){
	var task model.Task
	var res model.Response

	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&task)
	if err!=nil{
		res.Err="Error decode!"
		return
	}
	r.Body.Close()
	fmt.Println(task)
	err=dao.InsertTask(&task)
	if err!=nil{
		res.Err="Insert fail!"
		json.NewEncoder(w).Encode(res)
	}else{
		res.Success="Insert success!"
		json.NewEncoder(w).Encode(res)
	}

}

func UpdateTask(w http.ResponseWriter,r *http.Request){
	var task model.Task
	var res model.Response

	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&task)
	if err!=nil{
		fmt.Println("error decode update_task")
	}
	r.Body.Close()
	fmt.Println(task)

	err=dao.UpdateTask(&task)
	if err!=nil{
		res.Err="Update fail!"
		json.NewEncoder(w).Encode(res)
	}else{
		res.Success="Update success!"
		json.NewEncoder(w).Encode(res)
	}

}

func DeleteTask(w http.ResponseWriter, r *http.Request){
	var response model.Response

	param:=mux.Vars(r)
	id,_:=strconv.Atoi(param["id"])
	fmt.Println(id)

	err:=dao.DeleteTask(id)
	if err!=nil{
		response.Err="Delete fail!"
		json.NewEncoder(w).Encode(response)
	}else{
		response.Success="Delete success!"
		json.NewEncoder(w).Encode(response)
	}

}

func GetAllTasks(w http.ResponseWriter, r *http.Request){
	var tasks[] model.Task

	param:=mux.Vars(r)
	id,_:=strconv.Atoi(param["id"])
	fmt.Println(id)

	tasks=dao.GetAllTasks(id)
	
	json.NewEncoder(w).Encode(tasks)
	
}

func GetTask(w http.ResponseWriter, r *http.Request){
	var task model.Task

	param:=mux.Vars(r)
	id,_:=strconv.Atoi(param["id"])
	fmt.Println("gettask.dao",id)
	task=dao.GetTask(id)

	json.NewEncoder(w).Encode(task)
}

