package main

import "fmt"

func main(){
	cards := newDeck()
	cards = append(cards, "and Virgin Mojito")	
	newCards := deck{newCard(), "a deck card"}
	fmt.Println(cards)
	fmt.Println(newCards)	
	for i, card := range cards{
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "hello of diamonds"
}
