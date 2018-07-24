package main

func main(){
	cards := newDeckFromFile("my_file")
	cards.print()
	cards.shuffle()
}

