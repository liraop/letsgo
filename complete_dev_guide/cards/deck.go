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
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}
	return cards
}

// multiple returns - state how many elements
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {
	deckSize := len(d)
	rand.Seed(time.Now().UnixNano())
	for c, card := range d {
		randInd := rand.Intn(deckSize)
		randIndElement := d[randInd]
		d[randInd] = card
		d[c] = randIndElement
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(path string) {
	ioutil.WriteFile(path, []byte(d.toString()), 0666)
}

func newDeckFromFile(path string) deck {
	fileContentsBytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fileContentsStr := string(fileContentsBytes)
	splitDeck := strings.Split(fileContentsStr, ",")
	// option01
	// var newDeck deck = splitDeck
	// return newDeck
	// or option02
	return deck(splitDeck)
}

// d deck: d acts as value (receiver)
// for any deck (new type) variable
// to the function.
// in other words, as a This.
// must be at most 2 letters
// from the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
