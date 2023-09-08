package main

func main() {
	// ======== Printing Cards ========
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// ======== Save File ========
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// ======== Read File And Print It ========
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// ======== shuffle Cards And Print It ========
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
