package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// we create a custom data type 'deck' which has all functionalities of a slice
type deck []string

// we create a custom function that prints all values in this data type
// in this function we need to specify a receiver that will tell go that this function belongs to 'deck'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// a function to return a new slice of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// we are returning 2 things from this function, both are deck
func deal(d deck, hand_size int16) (deck, deck) {
	return d[:hand_size], d[hand_size:]
}

// a func to convert deck to string
func (d deck) toString() string {
	// var res_string string
	// for _, card := range d {
	// 	res_string += card
	// 	res_string += " "
	// }
	// return res_string

	// OR

	return strings.Join(d, ", ")
}

// a func to save our deck to local machine by saving it into a file
// error is a type of data type
// saveToFile also takes in a permissions as arguments
//
//	0666 > anyone can read and write
func (d deck) saveToFile(fileName string) error {
	// ioutil.WriteFile is deprecated
	// it used io/ioutil package
	// return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

// a function that turns saved text file to deck
func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	// bs is byteslice
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	tempString := string(bs)
	return strings.Split(tempString, ", ")
	// converting the above result to type deck is important, even though both are []string
	// this is because we can only use functions with type deck and not []string
	// ss := strings.Split(tempString, ", ")
	// return deck(ss)
	// naa checked it, newDeck.print() still runs without use type converting it
}

// a function to shuffle the deck of cards
func (d deck) shuffle() {
	n := len(d)
	for i := range d {
		ri := rand.Intn(n) // [0,n)
		d[i], d[ri] = d[ri], d[i]
	}

	// source := rand.NewSource(time.Now().UnixNano()) // returns time in nanoseconds > int64
	// r := rand.New(source)
	// ri := r.Intn(n - 1)
}
