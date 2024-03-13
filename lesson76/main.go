package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", Puffrider)
	http.ListenAndServe("0.0.0.0:80", nil)
}

func Puffrider(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, r.URL.Path)
	fmt.Println(r.URL.RawQuery)
	b, _ := io.ReadAll(r.Body)
	//fmt.Println(string(b))
	y := Request{}
	json.Unmarshal(b, &y)
	fmt.Println("Пришел запрос", y.FirstName, y.LastNme, "С номерами телефонов", y.PhoneNumbers)
}

type Request struct {
	FirstName    string `json:"firstName"`
	LastNme      string `json:"lastName"`
	PhoneNumbers []int  `json:"phoneNumbers"`
}
