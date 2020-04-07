package main


func main() {
	// creates a ne deck
	cards := newDeck()

	// shuffles the deck
	cards.shuffle()

	// prints the deck
	cards.print()

	// Save the deck to a file
	cards.saveToFile("cardsample")

	// load deck from file
	loadedDeck := newDeckFromFile("cardsample")

	// prints the loaded deck
	loadedDeck.print()
}

 