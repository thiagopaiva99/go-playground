package main

import "fmt"

func main() {
	users := []string{"Thiago", "Nicole", "Ariel", "Kevin"}
	pets := users[2:4]

	fmt.Println(pets)
}
