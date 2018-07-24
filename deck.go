package main

import (
	"fmt" 
	"strings"
	"io/ioutil"
)


type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() (string) {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error{
	return ioutil.WriteFile(fileName, []byte(d.toString()), 666) 
}
func deal (d deck, handsize int) (deck, deck){
	return d[:handsize],d[handsize:] 
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
