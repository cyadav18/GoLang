package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...))
	fmt.Println(sum1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println("--------------------------------------------------------------------------")
	p := person{
		first: "chethan",
		last:  "yadav",
		age:   23,
	}
	p.speak()
	fmt.Println("--------------------------------------------------------------------------")
	c := circle{
		radius: 22,
	}
	c.area()
	s := square{
		size: 23,
	}
	s.area()
	shapeInfo(s)
	shapeInfo(c)
	fmt.Println("--------------------------------------------------------------------------")
	//Question 6
	func() {
		fmt.Println("Any nomous func")
	}()
	func(a int) {
		for i := 1; i <= 10; i++ {
			fmt.Println(a, "*", i, "=", a*i)
		}
	}(2)
	fmt.Println("--------------------------------------------------------------------------")
	//Question 7
	add := func(lst ...int) int {
		total := 0
		for _, v := range lst {
			total += v
		}
		return total
	}
	fmt.Println(add([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...))
	fmt.Println("--------------------------------------------------------------------------")
	//Question 8
	add1 := addingSum()
	fmt.Println(add1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}...))
	fmt.Println("--------------------------------------------------------------------------")
	//Question 9
	fmt.Println(passingFunc(calculatePower, 2, 4))
	fmt.Println("--------------------------------------------------------------------------")
	//Question 10
	iterator1 := increment()
	fmt.Println(iterator1())
	fmt.Println(iterator1())
	fmt.Println(iterator1())
	fmt.Println(iterator1())
	fmt.Println(iterator1())
	fmt.Println(".............")
	iterator2 := increment()
	fmt.Println(iterator2())
	fmt.Println(iterator2())
	fmt.Println(iterator2())
	fmt.Println(iterator2())
	fmt.Println(iterator2())
}

//Question 1
func foo() int {
	fmt.Println("foo function called")
	return 4
}
func bar() (int, string) {
	fmt.Println("bar function called")
	return 5, "five"
}

//Question 2
func sum(lst ...int) int {
	total := 0
	for _, v := range lst {
		total += v
	}
	return total
}
func sum1(lst []int) int {
	total := 0
	for _, v := range lst {
		total += v
	}
	return total
}

//Question 3
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Person speaking :-", "name is", p.first, " ", p.last, " and age is ", p.age)
}

//Question 5
type circle struct {
	radius float64
}
type square struct {
	size float64
}

func (c circle) area() {
	fmt.Println("The area of circle is ", 3.142*c.radius*c.radius)
}

func (s square) area() {
	fmt.Println("The area of square is ", s.size*s.size)
}

type shape interface {
	area()
}

func shapeInfo(s shape) {
	switch s.(type) {
	case circle:
		s.area()
	case square:
		s.area()
	default:
		fmt.Println("No shape is selected")
	}
}

//Question 8

func addingSum() func(lst ...int) int {
	return func(lst ...int) int {
		total := 0
		for _, v := range lst {
			total += v
		}
		return total
	}
}

// Question 9

func calculatePower(num, pow int) int {
	if pow == 0 {
		return 1
	}
	pow--
	return num * calculatePower(num, pow)
}

func passingFunc(f func(num, pow int) int, nm, pw int) int {
	return f(nm, pw)
}

//Question 10

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
