package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", Cookie)
	http.HandleFunc("/2", Cookie2)

	http.ListenAndServe(":8080", nil)
}

// 传统设置法
func Cookie(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "heyguys",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	http.SetCookie(w, ck)

	ck2, err := r.Cookie("myCookie")
	if err != nil {
		// 错误表示 Cookie 不存在
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, ck2.Value)
}

// Header 设置法
func Cookie2(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hey guys",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	w.Header().Set("Set-Cookie", strings.Replace(ck.String(), " ", "%20", -1))

	ck2, err := r.Cookie("myCookie")
	if err != nil {
		// 错误表示 Cookie 不存在
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, ck2.Value)
}
