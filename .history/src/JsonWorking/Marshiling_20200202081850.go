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

//student obj
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
	p1 := Person{
		firstName:       "Chethan",
		lastName:        "Yadav",
		favoriteFlavour: []string{"blueberry", "cranberry", "strawberry"},
	}
	fmt.Println(p1)
}
