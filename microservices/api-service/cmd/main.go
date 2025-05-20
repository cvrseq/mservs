package main

import (
	"fmt"
	"log"
	"net/http"
	"services/api-service/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	handler.Init()

	router.HandleFunc("/employees", handler.GetEmployees).Methods("GET")
	router.HandleFunc("/employee/{id}", handler.GetEmployee).Methods("GET")
	router.HandleFunc("/employees", handler.AddEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", handler.UpdateEmployee).Methods("PUT")
	
	fmt.Println("Сервер запущен на порту 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}