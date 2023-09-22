package main

func main() {
	cards := newDeckFromFile("my_cards")

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	//cards.print()
	cards.shuffle()
}