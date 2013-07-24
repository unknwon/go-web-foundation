package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})

	// 使用函数作为 handler
	mux.HandleFunc("/bye", sayBye)

	// 静态文件
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My server: "+r.URL.String())
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye bye, this is version 2.")
}
