package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	user := person{firstName: "Thiago", lastName: "Galvani"}
	fmt.Println(user)
}
