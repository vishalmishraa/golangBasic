package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>this is my awesome website , hello vishal</h1>")
	} else if r.URL.Path == "/contacts" {
		fmt.Fprint(w, "To get in touch please send an email to <a href=\"mailto:vishalmishraa@gmail.com\">vsihal mishra</a>.")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
