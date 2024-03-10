package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// x := fildi{
	// 	Name:    "Vlad",
	// 	Age:     25,
	// 	Sex:     true,
	// 	Weight:  80.5,
	// 	Numbers: [5]int{1, 2, 3, 4, 5},
	// 	Races:   []string{"Negr", "Beliy", "Orange"},
	// 	Home: Adress{
	// 		City:    "Moscow",
	// 		Street:  "Arbat",
	// 		Bilding: 13,
	// 	},
	// 	Password: "1,4,6,4,3,",
	// 	Bonus:    5,
	// }
	// b, _ := json.MarshalIndent(x, "", "  ")
	// fmt.Println(string(b))
	var x fildi
	json.Unmarshal([]byte(y), &x)
	fmt.Printf("%+v\n", x)
}

type fildi struct {
	Name     string
	Age      int `json:"vozrast"`
	Sex      bool
	Weight   float64
	Numbers  [5]int
	Races    []string
	Home     Adress
	Password string
	Bonus    int `json:"bonus,omitempty"`
}
type Adress struct {
	City    string
	Street  string
	Bilding int
}

var y = `{
	"Name": "Vlad",
	"vozrast": 25,
	"Sex": true,
	"Weight": 80.5,
	"Numbers": [
	  1,
	  2,
	  3,
	  4,
	  5
	],
	"Races": [
	  "Negr",
	  "Beliy",
	  "Orange"
	],
	"Home": {
	  "City": "Moscow",
	  "Street": "Arbat",
	  "Bilding": 13
	},
	"bonus": 5,
	"Password":"1,2,3,4,5"
  }`
