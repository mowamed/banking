package app

import (
	"github.com/gorilla/mux"
	"github.com/mowamed/banking/domain"
	"github.com/mowamed/banking/service"
	"log"
	"net/http"
	"time"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// routes
	router.HandleFunc("/customers", ch.getAllCustomers)

	// server
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
