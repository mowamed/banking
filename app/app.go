package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	fmt.Fprintf(w, customerId)
}
