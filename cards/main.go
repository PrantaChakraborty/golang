package main

func main() {
	//var card string = "Ace of spades" // declare variable
	//card := "Ace of spades"   // when initializing a value with a new variable
	//card = "five of diamonds" // assigning another value
	//card := newCard() // call a function
	//fmt.Println(card)

	// slice
	//cards := newDeck()
	cards := newDeckFromFile("my_cards")
	cards.print()
	//looping through the cards
	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}
	//cards.print()
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")

}
