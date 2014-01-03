package main

import (
	"net/http"
)

type SingleHost struct {
	handler     http.Handler
	allowedHost string
}

func NewSingleHost(handler http.Handler, allowedHost string) *SingleHost {
	return &SingleHost{handler: handler, allowedHost: allowedHost}
}

func (s *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	if host == s.allowedHost {
		s.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success!"))
}

func main() {
	single := NewSingleHost(http.HandlerFunc(myHandler), "example.com")

	println("Listening on port 8080")
	http.ListenAndServe(":8080", single)
}
