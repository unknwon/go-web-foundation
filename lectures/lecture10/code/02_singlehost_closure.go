package main

import (
	"net/http"
)

func SingleHost(handler http.Handler, allowedHost string) http.Handler {
	ourFunc := func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		if host == allowedHost {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}
	}
	return http.HandlerFunc(ourFunc)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success!"))
}

func main() {
	single := SingleHost(http.HandlerFunc(myHandler), "example.com")

	println("Listening on port 8080")
	http.ListenAndServe(":8080", single)
}
