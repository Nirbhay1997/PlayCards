package main

func main() {
	// using a custom type created  in Deck.go
	cards := newDeckFromFile("Testfile")
	// adding an element to the slice
	//fmt.Println(cards.DeckToString())
	cards.print()
	// iterating a slice
	hand, remaining_cards := deal(cards, 5)
	hand.print()
	remaining_cards.print()

}
