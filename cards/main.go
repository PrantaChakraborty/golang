package main

import "fmt"

func main() {
	//var card string = "Ace of spades" // declare variable
	//card := "Ace of spades"   // when initializing a value with a new variable
	//card = "five of diamonds" // assigning another value
	//card := newCard() // call a function
	//fmt.Println(card)

	// slice
	cards := []string{"hello", "hi", newCard()}
	fmt.Println(cards)
	cards = append(cards, "Six of hearts")

	//	looping through the cards
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// function with return type
func newCard() string {
	return "Five of diamonds"
}
