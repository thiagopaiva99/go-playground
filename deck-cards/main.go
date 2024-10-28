package main

import "fmt"

func main() {
	cards := createDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	loadedDeck := newDeckFromFile("my_cards")
	fmt.Println(loadedDeck.toString())
}
