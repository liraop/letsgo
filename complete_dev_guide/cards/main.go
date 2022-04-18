package main

//import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//cards := newDeck()

	// append : returns new slice with
	// new element at the bottom
	//cards = append(cards, "Six of Spades")

	//hand, remainingDeck := deal(cards, 5)
	//
	//fmt.Println("Hand")
	//hand.print()
	//fmt.Println("Deck")
	//remainingDeck.print()
	// range: returns current
	// index and element
	//for i, card := range cards {
	//fmt.Println(i, card)
	//}

	cards := newDeckFromFile("test.txt")
	cards.print()
   cards.shuffle()
   cards.print()
}
