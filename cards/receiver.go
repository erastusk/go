package main

import "fmt"

func main() {
	cards := newDeck()
	// cards.print_cards()
	// Receivers - Declared a variable cards with the deck type. Now it has access to any 'methods' created for type deck.
	// can now use print_cards() method in deck.go
	// dealhand, remaining := deal(cards, 3)
	// dealhand.print_cards()
	// remaining.print_cards()
	// Type conversion
	// Type we want (value)
	//fmt.Println([]byte(cards.toString()))
	fmt.Println(cards.saveToFile("cardsfile"))
	fmt.Println(readFromFile("cardsfile"))
	cards.print_cards()
	cards = cards.shuffleDeck()
	cards.print_cards()

}

func newCard() string {
	return "Five of Diamonds"
}

// 1. Receivers are access like methods with a .
// 2. Functions with no receivers are called with args.
