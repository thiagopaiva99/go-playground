package main

import "fmt"

func main() {
	cards := createDeck()
	fmt.Println(cards.toString())
}
