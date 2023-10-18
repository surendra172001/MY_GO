package main

func main() {
	// var cards = createDeck()
	// cards.printDeck()

	// hand, remaining := deal(cards, 5)
	// fmt.Println(hand, remaining)

	// fmt.Println(cards.toString())

	// cards.saveFile("deck1")

	newCards := newDeckFromFile("Decks/deck1.txt")

	newCards.printDeck()

	newCards.shuffle()
	newCards.printDeck()
}
