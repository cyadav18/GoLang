package main

import "fmt"

type Hotdog int

var x Hotdog
var y int

func main() {
	var k bool
	k = primePredictor(31)
	x = 12
	y = int(x)
	fmt.Printf("the given number is prime %t \n", k)
	fmt.Printf("The type of x %T \n", x)
	fmt.Printf("The value of x is %d\n", x)
	fmt.Printf("The type of y %T \n", y)
	fmt.Printf("The value of y is %d\n", y)
}

func primePredictor(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	count := 0
	for j := 2; j <= num/2; j++ {
		if num%j == 0 {
			count++
		}
	}
	if count == 0 {
		return true
	}
	return false
}
