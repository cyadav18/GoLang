package main

import (
	"fmt"
)

type Person struct {
	firstName       string
	lastName        string
	favoriteFlavour []string
}

//Person exported
func main() {
	fmt.Println("Hello world")
	p1 := Person{
		firstName:       "Chethan",
		lastName:        "Yadav",
		favoriteFlavour: []string{"blueberry", "cranberry", "strawberry"},
	}
	fmt.Println(p1)
}
