package main

import "fmt"

type User struct {
	Name string
	Next *User
}

func main() {

	var list *User
	list = add(list, "Jonh")
	list = add(list, "Djas")
	list = add(list, "Dima")

	printList(list)

}

func printList(list *User) {

	for {
		if list == nil {
			break
		}

		fmt.Println(list.Name)
		list = list.Next
	}
}

func add(list *User, name string) *User {
	if list == nil {
		list = &User{
			Name: name,
		}
		return list
	}
	root := list
	for {
		if list.Next == nil {
			break
		}
		list = list.Next
	}
	list.Next = &User{
		Name: name,
	}
	return root
}
