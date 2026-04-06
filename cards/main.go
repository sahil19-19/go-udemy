package main

func main() {
	cards := newDeck()
	cards = append(cards, "King of hearts")
	// append doesnt modify the array but rather returns a new array

	cards.shuffle()

	// fmt.Println(cards)
	cards.print()

	// i, card are initialized every time in the for loop
	// cards.print()

	// hand, remaining_hand := deal(cards, 5)
	// fmt.Println(hand, remaining_hand)

	// fmt.Println(cards.toString())

	// saving deck to a file
	// cards.saveToFile("myCards")

	// getting back a deck from a saved file
	// newDeck := newDeckFromFile("my")
	// newDeck.print()
	// fmt.Println(newDeck)
}

func newFunc() string {
	return "Five of diamonds"
}
