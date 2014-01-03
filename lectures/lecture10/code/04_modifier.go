package main

import (
	"net/http"
	"net/http/httptest"
)

type ModifierMiddleware struct {
	handler http.Handler
}

func (m *ModifierMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := httptest.NewRecorder()
	// passing a ResponseRecorder instead of the original RW
	m.handler.ServeHTTP(rec, r)
	// after this finishes, we have the response recorded
	// and can modify it before copying it to the original RW

	// we copy the original headers first
	for k, v := range rec.Header() {
		w.Header()[k] = v
	}
	// and set an additional one
	w.Header().Set("X-We-Modified-This", "Yup")
	// only then the status code, as this call writes the headers as well
	w.WriteHeader(418)
	// the body hasn't been written yet, so we can prepend some data.
	w.Write([]byte("Middleware says hello again. "))
	// then write out the original body
	w.Write(rec.Body.Bytes())
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success!"))
}

func main() {
	mid := &ModifierMiddleware{http.HandlerFunc(myHandler)}

	println("Listening on port 8080")
	http.ListenAndServe(":8080", mid)
}
