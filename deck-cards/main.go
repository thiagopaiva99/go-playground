package main

func main() {
	cards := createDeck()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	// loadedDeck := newDeckFromFile("my_cards")
	// loadedDeck.print()

	newDeck := createDeck()
	newDeck.shuffle()
	newDeck.print()
}
