package main

// var card string = "Ace of Spades"
// card := "Ace of Spades"
//the := is only used for initialzion of the var you can use = to reassign it

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("my_cards")

	cards.print()
}
