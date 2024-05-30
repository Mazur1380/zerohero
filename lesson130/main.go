package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var data = `{"cart": ["Watermelon","Apple","Apple","Orange","Apple","Watermelon"]}`

func main() {

	var d interface{}
	err := json.Unmarshal([]byte(data), &d)
	if err != nil {
		log.Fatal(err)
	}

	m, ok := d.(map[string]interface{})
	if !ok {
		log.Fatal("invalid operation")
	}

	p, ok := m["cart"].([]interface{})
	if !ok {
		log.Fatal("invalid operation")
	}

	var prod = map[string]int{}

	for _, x := range p {
		p, ok := x.(string)
		if !ok {
			log.Fatal("not string")
		}
		prod[p] = prod[p] + 1
	}
	fmt.Println(prod)
}
