package main

import "fmt"

func main() {
	mySquare := square{
		sideLength: 100,
	}
	myTriangle := triangle{
		base:   200,
		height: 300,
	}
	printArea(mySquare)
	printArea(myTriangle)
}

type shape interface {
	getArea() float64
}

/*
	any struct with function `getArea() float64` is also a `shape`
	hence now, both triangle and square are also type shape
*/

func printArea(s shape) {
	fmt.Println("Area is", s.getArea())
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
