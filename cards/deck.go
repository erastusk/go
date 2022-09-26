package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck which is a slice of strings
type deck []string

func newDeck() deck {
	card := deck{}
	cardValues := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardSuits := []string{"Ace", "Two", "Three", "Four"}
	for _, Suit := range cardSuits {
		for _, Value := range cardValues {
			card = append(card, Suit+" of "+Value)
		}
	}
	return card
}

// func print_cards(d deck) {
// a function with a receiver type deck req not args and returns nothing.
func (d deck) print_cards() {
	for i, j := range d {
		fmt.Println(i, j)
	}

}

// a function that requires args of deck and int and returns 2 decks
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//[]byte -> string -> split into []string -> deck
	// could just return []string
	return deck(strings.Split(string(d), ","))
}

func (d deck) shuffleDeck() deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d

}
