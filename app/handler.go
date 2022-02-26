package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/dhruvbehl/bank/service"
)

// func greetHandler(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprint(w, "Hello world!!!")
// }

type CustomerHandler struct {
	service service.DefaultCustomerService
}

func (c *CustomerHandler) getAllCustomersHandler(w http.ResponseWriter, req *http.Request) {

	customers, _ := c.service.GetAllCustomer()
	if req.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

// func getCustomerByIdHandler(w http.ResponseWriter, req *http.Request) {
// 	id := mux.Vars(req)["customer_id"]
// 	for _, customer := range customers {
// 		if customer.Id == string(id) {
// 			if req.Header.Get("Content-Type") == "application/xml" {
// 				w.Header().Add("Content-Type", "application/xml")
// 				xml.NewEncoder(w).Encode(customer)
// 			} else {
// 				w.Header().Add("Content-Type", "application/json")
// 				json.NewEncoder(w).Encode(customer)
// 			}
// 			break
// 		}
// 	}
// }

// func createCustomerHandler(w http.ResponseWriter, req *http.Request) {
// 	var customerData CustomerDetails
// 	json.NewDecoder(req.Body).Decode(&customerData)
// 	for _, c := range customers {
// 		if c.Id == customerData.Id {
// 			json.NewEncoder(w).Encode(fmt.Sprintf("Customer with id [%v] already exists", customerData.Id))
// 			return
// 		}
// 	}
// 	customers = append(customers, customerData)
// 	json.NewEncoder(w).Encode(customers)
// }