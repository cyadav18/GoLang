package main

import "fmt"

func main() {
	var k bool
	k = primePredictor(31)
	fmt.Printf("the given number is prime %t", k)
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
