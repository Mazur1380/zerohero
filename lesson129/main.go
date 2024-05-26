package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var js = []byte(`{
"name": "Vlad",
"age":28,
"city":"SPB",
"hobbies":["Hookah","Sport","Programming"],
"address":{
		"street":"Kupchino",
		"building":34
}
}`)

func main() {
	var obj interface{}
	err := json.Unmarshal(js, &obj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(obj)

	m := obj.(map[string]interface{})
	v := m["age"]
	age := v.(float64)
	age = age + 1
	fmt.Println(age)

	h := m["hobbies"]
	hobbies := h.([]interface{})
	fmt.Println(len(hobbies))

	a := m["address"]
	address := a.(map[string]interface{})
	s := address["street"]
	street := s.(string)
	fmt.Println(street)
}
