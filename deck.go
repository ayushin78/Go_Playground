package main

import (
	"math/rand"
	"os"
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

func deal (d deck, handsize int) (deck, deck){
	return d[:handsize],d[handsize:] 
}

func newDeck() deck { 
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two","Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+ suit)
		}
	}
	
	return cards
}

func (d deck) saveToFile(fileName string) error{
	return ioutil.WriteFile(fileName, []byte(d.toString()), 666) 
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if(err != nil){
		fmt.Println("error : ", err)
		os.Exit(1);
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		j := rand.Intn(len(d))
		d[i], d[j] = d[j], d[i]
	}
}