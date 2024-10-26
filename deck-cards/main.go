package main

func main() {
	cards := createDeck()
	cards = append(cards, "Six of Spades")

	cards.print()
}
