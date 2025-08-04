package deck

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// NewDeck creates a new deck of cards
// It returns a deck containing 52 cards.
func NewDeck() deck {
	log.Println("Creating a new deck of cards")
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) Print() {
	for i, card := range d {
		println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func NewDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the erro and return a call to newDeck()
		// Option #2 - log the error and entirely quite the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss), nil
}

func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
