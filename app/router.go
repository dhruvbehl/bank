package app

import (
	"log"
	"net/http"
)

func Start() {
	router := http.NewServeMux()

	router.HandleFunc("/greet", greetHandler)
	router.HandleFunc("/getAllCustomers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}