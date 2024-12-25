package main

import "fmt"

type person struct {
	name string
	age  int
	exp  int
}

func (p person) Info() {
	fmt.Printf("my name %s, my age %d, my exp %d\n", p.name, p.age, p.exp)
}

type employee struct {
	role string
	exp  int
	person
}

func (e employee) Info() {
	fmt.Printf("my role %s, my exp %d\n", e.role, e.exp)
}

func main() {
	em := employee{
		role: "manager",
		person: person{
			name: "Anatoliy",
			age:  50,
			exp:  10,
		},
	}
	fmt.Printf("employee:%+v\n", em)
	em.person.Info()

}
