package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>this is my awesome website , hello vishal</h1>")
	} else if r.URL.Path == "/contacts" {
		fmt.Fprint(w, "To get in touch please send an email to <a href=\"mailto:vishalmishraa@gmail.com\">vsihal mishra</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>page not found:404</h1>")
	}

}

//hello is ...
func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/hello/:name", hello)

	// http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", router)
}
