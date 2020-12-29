package main

func main() {
	cards := newDeck()

	cards.shuffleDeck()
	cards.print()
	//hand, remainingDeck := deal(cards, 5) //first return value is gonna be = hand and the other one remainingDeck

	//cards.print() //Calling receiver from deck.go instead of writing a for loop
}
