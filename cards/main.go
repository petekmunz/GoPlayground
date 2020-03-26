package main

func main() {
	cards := newDeckFromFile("my_Cards")
	cards.shuffle()
	cards.print()
}
