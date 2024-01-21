package main

func main() {
	cards := deck{"test", newCard()}
	cards = append(cards, "eiei")

	cards.print()

}

func newCard() string {
	return "test 2"
}
