package main

import (
	"fmt"
	"math"
)

type hotdog int

var x hotdog
var y int
var number = []int{}
var exp = []int{}

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

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
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{1}
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, kb, mb/kb, gb/(mb))
	a = append(a, b...)
	fmt.Println(a)
	primefactor(513)
	fmt.Println(number, "\n", exp)
}

func primefactor(num int) {
	if primePredictor(num) {
		number = append(number, num)
		exp = append(exp, 1)
	} else {
		for i := 2; num != 1; i++ {
			if primePredictor(i) && num%i == 0 {
				number = append(number, i)
				exp = append(exp, getnumExpo(num, i))
				fst := float64(number[len(number)-1])
				snd := float64(exp[len(exp)-1])
				num = num / int(math.Pow(fst, snd))
			}
		}
	}
}

func getnumExpo(num1, num2 int) int {
	count := 0
	for num1 != 0 {
		if num1%num2 == 0 {
			count++
			num1 = num1 / num2
		} else {
			return count
		}
	}
	return -1
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
