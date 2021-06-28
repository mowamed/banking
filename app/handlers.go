package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/mowamed/banking/service"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, request *http.Request) {

	customers, _ := ch.service.GetAllCustomer()

	if request.Header.Get("Content-type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		log.Println(xml.NewEncoder(w).Encode(customers))
	} else {
		w.Header().Add("Content-type", "application/json")
		log.Println(json.NewEncoder(w).Encode(customers))
	}

}
