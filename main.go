package main

func main() {
	// cards := newDeck()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString(", "))

	// cards.saveToFile("my_cards.md")

	// cards := newDeckFromFile("my_cards.md")
	// cards.print()
	// fmt.Println(cards.toString(", "))

	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
}
