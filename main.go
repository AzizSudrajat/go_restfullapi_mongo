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
  route.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	route.HandleFunc("/api/books", createBook).Methods("POST")
	route.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	route.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	config := conn.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, route))

}
