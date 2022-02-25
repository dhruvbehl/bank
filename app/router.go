package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/getAllCustomers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}