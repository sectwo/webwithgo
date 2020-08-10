package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", indexpage)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func indexpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!!")
}
