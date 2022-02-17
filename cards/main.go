package main

import "fmt"

//static type language
//dynamic type dont care the type of the variables

//we can create variables outside the main func
//var sarasa int

func main() {
	/*ways of declaring a variable
	var card string = "Ace of spades"
	we only use the : to initialize the variable
	card := "Ace of Spades"
	card := newCard()

	there is no OOP in GO
	cards := []string{newCard(), newCard(), "sarasa"} //slice of type string

		for i, card := range cards {
		fmt.Println(i, card)
		cards := deck{newCard(), newCard(), "sarasa"} //slice of type string
	cards = append(cards, "adding new string")    //when you append Go wont modify the slice instead will create a new slice
	}
	//Here Im doing a type conversion from String to a slice of byte
	[]byte(card)
	hand, remainingCards := deal(cards, 5)

	hand.print()

	remainingCards.print()
	cards.saveToFile("sarasa")
	*/

	cards := newDeckFromFile("sarasa")
	cards.print()
	cards.shuffle()
	fmt.Println("--------------------------------")
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
