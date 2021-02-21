package main

import (
	"log"
	"net/http"

	"./config"
  "./controllers"
	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	route := mux.NewRouter()

	route.HandleFunc("/api/books", bookController.GetBooks).Methods("GET")
  route.HandleFunc("/api/books/{id}", bookController.GetBook).Methods("GET")
	route.HandleFunc("/api/books", bookController.CreateBook).Methods("POST")
	route.HandleFunc("/api/books/{id}", bookController.UpdateBook).Methods("PUT")
	route.HandleFunc("/api/books/{id}", bookController.DeleteBook).Methods("DELETE")

	config := conn.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, route))

}
