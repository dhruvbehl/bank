package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"

	"github.com/dhruvbehl/bank/dto"
	"github.com/dhruvbehl/bank/errors"
	"github.com/dhruvbehl/bank/service"
	"github.com/gorilla/mux"
)

// func greetHandler(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprint(w, "Hello world!!!")
// }

type CustomerHandler struct {
	service service.DefaultCustomerService
}

func writeResponse(w http.ResponseWriter, req *http.Request, code int, data interface{}) {
	switch req.Header.Get("Content-Type") {
	case "application/xml":
		w.Header().Add("Content-Type", "application/xml")
		w.WriteHeader(code)
		if err := xml.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	default:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	}
}

func (c *CustomerHandler) getAllCustomersHandler(w http.ResponseWriter, req *http.Request) {
	status := req.URL.Query()["status"]
	var customers []dto.CustomerResponse
	var err *errors.AppError
	if len(status) < 1 {
		customers, err = c.service.GetAllCustomer()
	} else {
		if strings.EqualFold(status[0], "active") ||  strings.EqualFold(status[0], "inactive") {
			customers, err = c.service.GetCustomerByStatus(status[0])
		} else {
			err := errors.NewBadRequestError("query parameter status only accepts [active | inactive]")
			writeResponse(w, req, err.Code, err.Message)
		}
	}
	if err != nil {
		writeResponse(w, req, err.Code, err)
	} else {
		writeResponse(w, req, http.StatusOK, customers)
	}
}

func (c *CustomerHandler) getCustomerByIdHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["customer_id"]
	customer, err := c.service.GetCustomerById(id)
	if err != nil {
		writeResponse(w, req, err.Code, err)
	} else {
		writeResponse(w, req, http.StatusOK, customer)
	}
}

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