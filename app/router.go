package app

import (
	"log"
	"net/http"

	"github.com/dhruvbehl/bank/domain"
	"github.com/dhruvbehl/bank/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/getAllCustomers", ch.getAllCustomersHandler).Methods(http.MethodGet)
	// router.HandleFunc("/getCustomerById/{customer_id:[0-9]+}", getCustomerByIdHandler).Methods(http.MethodGet)
	// router.HandleFunc("/customer", createCustomerHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}