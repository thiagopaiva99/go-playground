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

func main() {
	user1 := person{firstName: "Thiago", lastName: "Galvani", contact: contactInfo{email: "thiago@example.com", zipCode: 12345}}
	fmt.Println(user1)

	var user2 person
	user2.firstName = "John"
	user2.lastName = "Doe"
	user2.contact.email = "john@example.com"
	user2.contact.zipCode = 67890
	fmt.Println(user2)
}
