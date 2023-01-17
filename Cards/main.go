package main
//import "fmt"

func main(){
	cards:= newDeck()
	cards.shuffle()
	cards.print()
}

func NewCard() string {
	return "Ace of hearts."
}