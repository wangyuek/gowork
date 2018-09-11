package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index(w http.ResponseWriter, r *http.Request) is the default called function
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello,%q", html.EscapeString(r.URL.Path))
}
