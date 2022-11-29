package cards

import (
	mydeck "practice/cards/deck"
)

func Cards() {
	cards := mydeck.NewDeck()

	/*
		error_ := cards.SaveToFile("mycards")
		if error_ != nil {
			fmt.Println("Error occured while writing to a file")
		} else {
			fmt.Println("All good!!!")
		}
		card1, card2 := cards.Deal(5)
		card1.Print()
		fmt.Printf("\n")
		card2.Print()
		fmt.Println("Reading from the file:")
		cards_ := NewDeckFromFile("mycards")
		cards_.Print()
	*/
	cards.Shuffle()
	cards.Print()
}
