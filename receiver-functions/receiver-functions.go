package main

import "fmt"

func main() {
	users := user{"Thiago", "Nicole", "Ariel", "Kevin"}
	users.print()

	filteredUsers := users.filterByName("Thiago")
	fmt.Println(filteredUsers)
}

type user []string

func (u user) print() {
	for i, user := range u {
		fmt.Println(i, user)
	}
}

func (u user) filterByName(name string) user {
	var filteredUser user
	for _, user := range u {
		if user == name {
			filteredUser = append(filteredUser, user)
		}
	}
	return filteredUser
}
