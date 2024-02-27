package main

import "fmt"

func main() {
	card := newCard()
	cards := deck{"new card", newCard()}

	cards = append(cards, "Six of Spades")

	for i, cards := range cards {
		fmt.Println(i, cards)
	}

	fmt.Println(cards)
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
