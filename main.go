package main

import (
	"github.com/gorilla/mux"
	customer "goproject/controller"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/customers", customer.GetCustomer).Methods("GET")
	http.Handle("/", r)
	println("Connected port :8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}
