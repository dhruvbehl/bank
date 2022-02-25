package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello world!!!")
}

func getAllCustomers(w http.ResponseWriter, req *http.Request) {
	customers := []CustomerDetails{
		{"Dhruv Behl", "Lucknow", "95033xxxxx"},
		{"Aditi Behl", "Lucknow", "91234xxxxx"},
	}

	if req.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}