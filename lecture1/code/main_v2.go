package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	log.Println("Starting server... v2")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye this is version 2!"))
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello v2, the request URL is: " + r.URL.String()))
}
