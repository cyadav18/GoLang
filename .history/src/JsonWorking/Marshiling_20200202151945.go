package main

import (
	"fmt"
	"strings"
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
		dorrNo:     "No 31",
		firstLine:  "5th Main, 7th Cross, SVG Nagar",
		secondLine: "Moodalapalya",
		state:      "Karnataka",
		district:   "Bangalore",
		pin:        "560072",
	}

	student := Student{
		person:  person,
		address: address,
		rollNo:  21,
	}
	student.printStudent()
}
func (a Address) printAddress() string {
	return ("Address:" + a.dorrNo + ", " + a.firstLine + ", " + a.secondLine + ", " + a.state + ", " + a.district + ", " + a.pin)
}

func (p Person) printPerson() string {
	return ("Person :" + p.firstName + " " + p.lastName + " icecream flavors [" + strings.Join(p.favoriteFlavour, `,`) + "]")
}

func (s Student) printStudent() string {
	return ("Student: " + s.person.printPerson + " roll no " + fmt.Sprint(s.rollNo) + " " + s.address.printAddress)
}
