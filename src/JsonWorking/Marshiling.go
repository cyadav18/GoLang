package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

//Person exported
type Person struct {
	FirstName       string
	LastName        string
	FavoriteFlavour []string
}

//Student obj
type student struct {
	Person  Person
	RollNo  int
	Address Address
}

//Address obj
type Address struct {
	DorrNo     string
	FirstLine  string
	SecondLine string
	State      string
	District   string
	Pin        string
}

func main() {
	fmt.Println("Hello world")
	person := Person{
		FirstName:       "Chethan",
		LastName:        "Yadav",
		FavoriteFlavour: []string{"blueberry", "cranberry", "strawberry"},
	}
	address := Address{
		DorrNo:     "No 31",
		FirstLine:  "5th Main, 7th Cross, SVG Nagar",
		SecondLine: "Moodalapalya",
		State:      "Karnataka",
		District:   "Bangalore",
		Pin:        "560072",
	}

	st := student{
		Person:  person,
		Address: address,
		RollNo:  21,
	}
	var1 := []student{st}
	bs, err := json.Marshal(var1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
func (a Address) printAddress() string {
	return ("Address:" + a.DorrNo + ", " + a.FirstLine + ", " + a.SecondLine + ", " + a.State + ", " + a.District + " - " + a.Pin)
}

func (p Person) printPerson() string {
	return ("Person :" + p.FirstName + " " + p.LastName + " icecream flavors [" + strings.Join(p.FavoriteFlavour, `,`) + "]")
}

func (s student) printStudent() string {
	return ("Student: " + s.Person.printPerson() + ", Rolno :" + fmt.Sprint(s.RollNo) + ", " + s.Address.printAddress())
}
