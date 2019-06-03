package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/handlers"
	"TODOLIST/routes"
)

func main() {
	port := ":8080"
	fmt.Println("Starting Server at ", port)
	// log.Fatal(http.ListenAndServe(port, handlers.CORS()(routes.LawRouter())))
	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD","DELETE", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(routes.Router())))
	// log.Fatal(http.ListenAndServe(port, routes.LawRouter()))
	
}