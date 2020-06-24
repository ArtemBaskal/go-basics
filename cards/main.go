package main

func main() {
	newDeck().saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
