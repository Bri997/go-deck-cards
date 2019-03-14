package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	//the := is only used for initialzion of the var you can use = to reassign it

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
