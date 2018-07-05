package main

import "fmt"

type deck []string


func (this deck) print() {
	for i, card := range this {
		fmt.Println(i, card)
	}
}

func newDeck() deck { 
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two","Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+ value)
		}
	}
	
	return cards
}
