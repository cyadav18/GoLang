package main

import (
	"fmt"
)

func main() {
	var x [10]int

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t %c\n", i, i)
	}
	for i := 0; i < len(x); i++ {
		x[i] = (i + 1) * 10
	}
	var y = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < len(y); i++ {
		fmt.Println(y[i])
	}
	z := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y = append(z, y...)
	fmt.Println("Hello")
	fmt.Println(x)
	fmt.Println(y)
	for i, v := range y {
		fmt.Printf("%d \t %d\n", i, v)
	}
}
