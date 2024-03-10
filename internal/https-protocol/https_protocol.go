package httpsprotocol

import (
	"fmt"
	"net/http"
)

func handle_foo(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println("header key: ", k, "header value: ", v)
	}
	w.Write([]byte("Hello, World!"))
}

func handle_login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
	} else {
		fmt.Println("Invalid request method.")
	}
}

func Http_protocol() {
	http.HandleFunc("/foo", handle_foo)
	http.HandleFunc("/login", handle_login)
	http.ListenAndServe(":8080", nil)
}