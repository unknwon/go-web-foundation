package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("go-web").Parse(`
Package name: {{.Name}}
Number of functions: {{.NumFuncs}}
Number of variables: {{.NumVars}}`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}
		err = tmpl.Execute(w, &Package{
			Name:     "go-web",
			NumFuncs: 12,
			NumVars:  1200,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
