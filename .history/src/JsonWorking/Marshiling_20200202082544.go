package main

import (
	"fmt"
)

//Person exported
type Person struct {
	firstName       string
	lastName        string
	favoriteFlavour []string
}

//Student obj
type Student struct {
	person  Person
	rollNo  int
	address Address
}

//Address obj
type Address struct {
	dorrNo     string
	firstLine  string
	secondLine string
	state      string
	district   string
	pin        string
}

func main() {
	fmt.Println("Hello world")
	person := Person{
		firstName:       "Chethan",
		lastName:        "Yadav",
		favoriteFlavour: []string{"blueberry", "cranberry", "strawberry"},
	}
	address := Address{
		dorrNo:     "#31",
		firstLine:  "5th Main, 7th Cross, SVG Nagar",
		secondLine: "Moodalapalya",
		state:      "Karnataka",
		district:   "Bangalore",
		pin:        "560072",
	}

	student := Student{
		person:  person,
		address: address,
		rollNo:21
	}
	fmt.Println(person, address, student)
}
