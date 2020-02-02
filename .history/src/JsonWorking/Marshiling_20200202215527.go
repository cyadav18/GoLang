package main

import (
	"encoding/json"
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
	Person  Person
	RollNo  int
	Address Address
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
		Person:  person,
		Address: address,
		RollNo:  21,
	}
	var1 := []Student{student}
	bs, err := json.Marshal(var1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
func (a Address) printAddress() string {
	return ("Address:" + a.dorrNo + ", " + a.firstLine + ", " + a.secondLine + ", " + a.state + ", " + a.district + " - " + a.pin)
}

func (p Person) printPerson() string {
	return ("Person :" + p.firstName + " " + p.lastName + " icecream flavors [" + strings.Join(p.favoriteFlavour, `,`) + "]")
}

func (s Student) printStudent() string {
	return ("Student: " + s.Person.printPerson() + ", Rolno :" + fmt.Sprint(s.RollNo) + ", " + s.Address.printAddress())
}
