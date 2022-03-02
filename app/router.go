package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhruvbehl/bank/domain"
	"github.com/dhruvbehl/bank/service"
	"github.com/gorilla/mux"
)

func Start(envVar domain.Environment) {
	router := mux.NewRouter()

	// ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb(envVar))}
	router.HandleFunc("/getAllCustomers", ch.getAllCustomersHandler).Methods(http.MethodGet)
	router.HandleFunc("/getCustomerById/{customer_id:[0-9]+}", ch.getCustomerByIdHandler).Methods(http.MethodGet)
	// router.HandleFunc("/customer", createCustomerHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v",envVar.HOST, envVar.PORT), router))
}