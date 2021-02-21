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
	r := mux.NewRouter()

	r.HandleFunc("/api/books", bookController.GetBooks).Methods("GET")

	config := conn.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
