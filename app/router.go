package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/greet", greetHandler).Methods(http.MethodGet)

	router.HandleFunc("/getAllCustomers", getAllCustomersHandler).Methods(http.MethodGet)
	router.HandleFunc("/getCustomerById/{customer_id:[0-9]+}", getCustomerByIdHandler).Methods(http.MethodGet)
	router.HandleFunc("/customer", createCustomerHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}