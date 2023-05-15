package main

import "fmt"

// create a new type of 'deck'
// which is a slice of string

type deck []string

// function to generate a list of deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Club"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
