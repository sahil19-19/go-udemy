package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed()
	for i := 0; i < 10; i++ {
		x := rand.Intn(2)
		fmt.Println(x)
	}
}
