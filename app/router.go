package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dhruvbehl/bank/domain"
	"github.com/dhruvbehl/bank/service"
	"github.com/gorilla/mux"
)

func Start() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if host == "" || port == "" {
		panic("Application needs HOST and PORT as environment variables")
	}
	router := mux.NewRouter()

	// ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/getAllCustomers", ch.getAllCustomersHandler).Methods(http.MethodGet)
	router.HandleFunc("/getCustomerById/{customer_id:[0-9]+}", ch.getCustomerByIdHandler).Methods(http.MethodGet)
	// router.HandleFunc("/customer", createCustomerHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), router))
}