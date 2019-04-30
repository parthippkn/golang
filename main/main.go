package main

import (
	"fmt"
	"golang/util"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<H2>Hello Go Server</H2>")
}

func main() {
	util.EqualsIgnoreCase("Hello")

	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	//http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", mux)
}
