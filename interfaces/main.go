package main

import "fmt"

type bot interface {
	getGreeting() string

	// if getGreeting accepted arguments of type int
	// getGreeting(int) string
	// note : we dont mention variables names

}

// bot is an interface type, and we cant create a value out of interface type
// interfaces are implicit

/*

 */

type englishBot struct {
}

type spanishBot struct {
}

func (eb englishBot) getGreeting() string {
	// some custom logic
	return "hello world"
}

// since we are not making use of 'sb' and 'eb' inside our function we can remove it
func (spanishBot) getGreeting() string {
	// some custom logic
	return "hello"
}

// we see that printGreeting function doesnt do anything eb or sb specific
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// in go function overloading isnt allowed, so we cant declare 2 functions with same name
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// we create a functions that takes a `bot` argument, thus can be called
// by anything that satifies to be of type `bot`
// a single function to work for both eb and sb
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
