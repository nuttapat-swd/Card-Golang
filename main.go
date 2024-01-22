package main

func main() {
	// cards := newDeck()

	deck := newDeckFromFile("test.txt")
	deck.print()

}
