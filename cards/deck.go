package main

import (
	"fmt"
	"os"
	"strings"
)

// create a new type of 'deck'
// which is a slice of string

type deck []string

// function to generate a list of deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Club"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver
// a deck type receiver to print list of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// function to return multiple values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // slicing
}

// receiver function
// for use deck.toString()
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// function to write on file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}
