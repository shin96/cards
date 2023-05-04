package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of Diamonds" // another way of creating a variable
	// fmt.Println(newDeck().toString())
	// newDeck().saveToFile("myCards.txt")
	for _, card := range newDeckFromFile("yCards.txt") {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}