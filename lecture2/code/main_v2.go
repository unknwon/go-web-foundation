package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd: %v", err)
	}
	log.Print("Work directory:", wd)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// tmpl := template.Must(template.ParseFiles(filepath.Join(wd, "main_v2.tmpl")))
		tmpl, err := template.ParseFiles(filepath.Join(wd, "main_v2.tmpl"))
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
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
