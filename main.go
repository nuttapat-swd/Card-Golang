package main

func main() {
	cards := newDeck()

	hands, remainingCards := deal(cards, 5)
	hands.print()
	remainingCards.print()
}
