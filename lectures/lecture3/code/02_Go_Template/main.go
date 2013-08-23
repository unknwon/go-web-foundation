package main

import (
	"html/template"
	"os"
)

func main() {
	t := template.New("new tpl")
	t, _ = t.Parse("{{.Name}}")
	var user struct{ Name string }
	user.Name = "Joe"
	t.Execute(os.Stdout, user)
	println()
}
