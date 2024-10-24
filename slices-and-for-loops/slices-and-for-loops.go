package main

import "fmt"

func main() {
	users := []string{"Thiago", "Nicole"}
	users = append(users, "Ariel")
	users = append(users, "Kevin")

	for i, user := range users {
		fmt.Println(i, user)
	}
}
