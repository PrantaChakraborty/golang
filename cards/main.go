package main

func main() {
	//var card string = "Ace of spades" // declare variable
	//card := "Ace of spades"   // when initializing a value with a new variable
	//card = "five of diamonds" // assigning another value
	//card := newCard() // call a function
	//fmt.Println(card)

	// slice
	cards := newDeck()

	//	looping through the cards
	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}
	cards.print()
}
