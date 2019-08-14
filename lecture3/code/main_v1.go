package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd: %v", err)
	}
	log.Print("Work directory:", wd)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("main_v1.tmpl").Funcs(template.FuncMap{
			"Add": func(num1, num2 int) int {
				return num1 + num2
			},
			"Subtract": func(num1, num2 int) int {
				return num1 - num2
			},
			"Multiple": func(num1, num2 int) int {
				return num1 * num2
			},
			"Divide": func(num1, num2 int) int {
				return num1 / num2
			},
			"Str2html": func(str string) template.HTML {
				return template.HTML(str)
			},
			"Date": func() time.Time {
				return time.Now()
			},
			"Square": func(num int) int {
				return num * num
			},
		}).ParseFiles(filepath.Join(wd, "main_v1.tmpl"))
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
			return
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"Link":  `<a href="https://gogs.io">Gogs Website</a>`,
			"Num1":  15,
			"Num2":  1024,
			"Slice": []string{"foo", "bar"},
			"Array": [2]string{"hello", "world"},
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
