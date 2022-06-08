package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my bad page!<h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch with <a href=\"mailto:alisher1999@gmail.com\">alisher1999@gmail.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404 not FOUND:(<h1><p>Please email us</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
