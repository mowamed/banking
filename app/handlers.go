package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, request *http.Request) {
	log.Println(fmt.Fprint(w, "hello world!"))
}

func getAllCustomers(w http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{"Mohamed Bakus", "Accra", "00225"},
		{"Dexter Ouatt", "Accra", "00225"},
		{"Momo bak", "Accra", "00225"},
	}

	if request.Header.Get("Content-type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		log.Println(xml.NewEncoder(w).Encode(customers))
	} else {
		w.Header().Add("Content-type", "application/json")
		log.Println(json.NewEncoder(w).Encode(customers))
	}

}
