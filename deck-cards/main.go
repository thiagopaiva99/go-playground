package main

import "fmt"

func main() {
	cards := createDeck()
	hand, remainingCards := deal(cards, 5)
	fmt.Println(hand)
	fmt.Println(remainingCards)
}
