package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", parseForm)
	http.ListenAndServe(":8080", nil)
}

const tpl = `
<html>
	<head>
		<title>hello world</title>
	</head>
	<body>
		<form method="post" action="/">
			Username: <input type="text" name="uname">
			Password: <input type="password" name="pwd">
			<button type="submit">Submit</button>
		</form>
	</body>
</html>
`

func parseForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.New("hey")
		t.Parse(tpl)
		t.Execute(w, nil)
	} else {
		//r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println(r.FormValue("uname"))
	}
}
