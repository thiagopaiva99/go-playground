package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	user1 := person{firstName: "Thiago", lastName: "Galvani"}
	fmt.Println(user1)

	var user2 person
	user2.firstName = "John"
	user2.lastName = "Doe"
	fmt.Println(user2)
}
