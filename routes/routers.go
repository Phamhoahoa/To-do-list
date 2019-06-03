package routes

import(
	"net/http"
	"github.com/gorilla/mux"
	"TODOLIST/controller"
)

func Router() http.Handler{
	router:=mux.NewRouter()
	router.HandleFunc("/api/user/register",controller.Register).Methods("POST")
	router.HandleFunc("/api/user/login",controller.Login).Methods("POST")

	router.HandleFunc("/api/task/insert",controller.InsertTask).Methods("POST")
	router.HandleFunc("/api/task/update",controller.UpdateTask).Methods("POST")
	router.HandleFunc("/api/task/delete/{id}",controller.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/task/getalltasks/{id}",controller.GetAllTasks).Methods("GET")
	router.HandleFunc("/api/task/gettask/{id}",controller.GetTask).Methods("GET")
	return router
}
