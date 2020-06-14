package main

import (
	"github.com/gorilla/mux"
	customer "goproject/controller"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	// routing endpoint to customer controller function
	router.HandleFunc("/customers", customer.GetCustomer).Methods("GET")

	http.Handle("/", router)

	println("Connected port :8000")
	log.Fatal(http.ListenAndServe(":8000",router))
}
