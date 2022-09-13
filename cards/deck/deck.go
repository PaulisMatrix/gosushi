package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck which is a string slice.
type deck []string

//receiver function which acts on the deck type we defined above
//copy of deck type is passed
func (d deck) Print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}
func (d deck) Deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) ToString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func NewDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

/*
Using struct to house suit and its value.
type mydeck struct {
	suit  string
	value string
}

func main() {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	var d []mydeck

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			d = append(d, mydeck{suit: suits, value: values})
		}
	}
	fmt.Println(d)

}*/
