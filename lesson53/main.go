package main

import (
	"fmt"
	"net/http"
)

func main() {

	x := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, "Hello world <a href=\"/pow\">pow</a>")

	}
	http.HandleFunc("/", x)
	http.HandleFunc("/pow", pow)
	http.ListenAndServe(":80", nil)

}
func pow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "Privet psi <a href=\"/\"> Nazad</a>")

}
