package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName   string
	lastName    string
	contact     contactInfo
	contactInfo // two ways, this makes the property name contactInfo as well
}

// we can write receiver functions for struct as well
// creates a copy for the struct
func (p Person) print() {
	fmt.Printf("%+v", p)
}

// whenever we pass a variable to a function it is passed by value, hence a new variable is created in memory
// the variable that we pass to this function is copied into p and all changes take place on p
// thus nothing is updated
// to solve this we pass values by address
func (p *Person) updateFirstName(newName string) {
	// (*p).firstName = newName
	p.firstName = newName
	/*
		Go automatically handles calling:
		even though method uses *Person
	*/
}

func main() {
	// // var person Person

	// // 1.
	// // sahil := Person{"sahil", "verma"}

	// // 2.
	// sahil := Person{
	// 	firstName: "sahil",
	// 	lastName:  "verma",
	// }

	// // 3. declaration
	// var alex Person // zero values assigned for all keys

	// fmt.Println(sahil)
	// fmt.Println(alex)

	// fmt.Printf("%+v \n", sahil)
	// // %+v prints out all the the field names and their values

	// alex.firstName = "alex"
	// fmt.Printf("%+v", alex)
	jim := Person{
		firstName: "jimmy",
		lastName:  "snow",
		contact: contactInfo{
			email:   "jimmy@gmail.com",
			zipCode: 12345,
		},
		contactInfo: contactInfo{
			email:   "newjimmy@gmail.com",
			zipCode: 12345,
		},
	}

	// fmt.Printf("%+v", jim)
	// jim.updateFirstName("John")
	// this wont change the value of jim

	jimPointer := &jim // store the address of jim in jimPointer
	jimPointer.updateFirstName("John")

	jim.updateFirstName("Johny")
	// this works correctly and updates the names to Johny
	// requires less code

	jim.print()
}
