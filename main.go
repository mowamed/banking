package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greet)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, request *http.Request) {
	log.Println(fmt.Fprint(w, "hello world!"))
}
