package main

import "fmt"

func main() {
	cards := newDeck()

	if cards.saveToFile("test.txt") != nil {
		fmt.Println("eieieieieiei")
	}
}
