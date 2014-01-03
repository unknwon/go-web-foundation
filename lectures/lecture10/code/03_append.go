package main

import (
	"net/http"
)

type AppendMiddleware struct {
	handler http.Handler
}

func (a *AppendMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
	w.Write([]byte("<!-- Middleware says hello! -->"))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success!"))
}

func main() {
	mid := &AppendMiddleware{http.HandlerFunc(myHandler)}

	println("Listening on port 8080")
	http.ListenAndServe(":8080", mid)
}
