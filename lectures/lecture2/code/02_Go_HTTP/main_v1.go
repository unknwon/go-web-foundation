package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// 设置访问路由
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/bye", sayBye)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version 1.")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye bye, this is version 1.")
}
