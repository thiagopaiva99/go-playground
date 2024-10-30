package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	user1 := person{firstName: "Thiago", lastName: "Galvani", contact: contactInfo{email: "thiago@example.com", zipCode: 12345}}
	user1.print()

	var user2 person
	user2.firstName = "John"
	user2.lastName = "Doe"
	user2.contact.email = "john@example.com"
	user2.contact.zipCode = 67890
	user2.print()

	user2.updateName("Jimmy")
	user2.print()
}
