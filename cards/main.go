package main

import "fmt"

func main() {
	//var card string = "Ace of spades" // declare variable
	//card := "Ace of spades"   // when initializing a value with a new variable
	//card = "five of diamonds" // assigning another value
	card := newCard() // call a function
	fmt.Println(card)
}

// function with return type
func newCard() string {
	return "Five of diamonds"
}
