package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	suits := [] string {"Diamonds", "Clubs", "Hearts", "Spades"}
	values := [] string {"Ace", "Two", "Three", "Four",}

	for _, suit:= range suits {
		for _, value := range values {
            cards = append(cards, value + " of " + suit)
        }
	}
    
	return cards
}

func (d deck) print() {
    for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteCards, err := ioutil.ReadFile(fileName)
	if err!= nil {
        fmt.Println("error:", err)
		os.Exit(1)
    }
	
	s := strings.Split(string(byteCards), ",")
	return deck(s)

}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
        j := r.Intn(len(d) - 1)
        d[i], d[j] = d[j], d[i]
    }
    return d
}
