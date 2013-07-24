package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8000",
		Handler:     &myHandler{},
		ReadTimeout: time.Second * 5, // 超时设置
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = sayHello
	mux["/bye"] = sayBye

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version 3.")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye bye, this is version 3.")
}
