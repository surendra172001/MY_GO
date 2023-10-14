package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Deck []string

func createDeck() Deck {
	var cards = Deck{}

	var suits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	var values = []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d Deck) addNewCard(newCard string) {
	d = append(d, newCard)
}

func (d Deck) printDeck() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) toByte() []byte {
	return []byte(d.toString())
}

func (d Deck) saveFile(fileName string) error {
	var root_loc, _ = os.Getwd()
	return os.WriteFile(filepath.Join(root_loc, "Decks", fileName+".txt"), d.toByte(), 0644)
}

func newDeckFromFile(fileName string) Deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println(string(data))

	data := string(bs)

	return Deck(strings.Split(data, ","))
}

func (d Deck) shuffle() {
	sv := int64(time.Now().Unix())
	// fmt.Printf("%T", int64(sv))
	rand.Seed(sv)
	// fmt.Println(rand.Intn(len(d) - 1))
	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
