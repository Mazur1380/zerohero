package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", Handler)
	fmt.Println("Вызываем сервер")
	http.ListenAndServe(":80", nil)
}
func Handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Header)
	//fmt.Println(r.Method)
	if r.Method == "GET" {
		fmt.Fprint(w, x)
		return
	}
	if r.Method == "POST" {
		x = x + 1
		fmt.Fprint(w, "Увеличили х")
		return
	}

}

var x int
