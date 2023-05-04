package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of Diamonds" // another way of creating a variable
	card := newCard()
	cards := [] string { "five of hearts", newCard() }
	cards = append(cards, "joker")
	
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}