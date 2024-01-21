package main

import "fmt"

// Create a new type of 'deck'
// Which is slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"ดอกจิก", "หลามตัด", "โพธ์แดง", "โพธ์ดำ"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
