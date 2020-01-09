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

type person struct {
	firstName       string
	lastName        string
	favoriteFlavour []string
}
type vehicle struct {
	door string
	string
	breaks bool
}

type truck struct {
	name string
	vehicle
	fourWheeler bool
}

type sedan struct {
	name        string
	vehicleSpec vehicle
	luxury      bool
}

func main() {
	//assignment4()
	assignment5()
}
func assignment5() {
	fmt.Println("__________________________________________________________")
	fmt.Println("Question1")
	p1 := person{
		firstName:       "Chethan",
		lastName:        "Yadav",
		favoriteFlavour: []string{"Venila", "Blueberry", "Chocolate"},
	}
	p2 := person{
		firstName:       "Vijethan",
		lastName:        "Yadav",
		favoriteFlavour: []string{"Venila", "Blueberry", "Chocolate"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("first name :", p1.firstName, "\nlast name :", p1.lastName, "\nfavorite icecreams :", p1.favoriteFlavour)
	fmt.Println("first name :", p2.firstName, "\nlast name :", p2.lastName, "\nfavorite icecreams :", p2.favoriteFlavour)
	fmt.Println("__________________________________________________________")
	fmt.Println("Question2")
	map1 := map[string]person{}
	map1[p1.firstName] = p1
	map1[p2.lastName] = p2
	fmt.Println(map1)
	fmt.Println("__________________________________________________________")
	fmt.Println("Question3")
	v1 := vehicle{
		door:   "normal door",
		string: "blue",
		breaks: true,
	}
	v2 := vehicle{
		door:   "extrodinary door",
		string: "red",
		breaks: true,
	}
	trk := truck{
		name: "Lorry",
		vehicle: vehicle{
			door:   "old door",
			string: "yellow",
			breaks: true,
		},
		fourWheeler:true
	}
	sdn := sedan{
		name:        "sedan",
		vehicleSpec: v2,
		luxury:true
	}
	fmt.Println(v1.door, v1.string)
	fmt.Println(v2)
	fmt.Println(trk)
	fmt.Println(sdn)
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

func assignment4() {
	// Question 1
	fmt.Println("________________________________________________________________________")
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)
	for i, v := range arr {
		fmt.Println("index = ", i, "range =", v)
	}
	fmt.Printf("Type of array is %T \n", arr)
	fmt.Println("________________________________________________________________________")
	//Question 2
	var arr1 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range arr1 {
		fmt.Println("index = ", i, "range =", v)
	}
	fmt.Println("The length of the array arr1 is = ", len(arr1))
	fmt.Println(arr1)
	fmt.Printf("Type of array is %T \n", arr1)
	fmt.Println("________________________________________________________________________")
	//Question 3
	var arr2 []int = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(arr2[:5])
	fmt.Println(arr2[5:])
	fmt.Println(arr2[2:7])
	fmt.Println(arr2[1:6])
	fmt.Println("________________________________________________________________________")
	//Question 4
	fmt.Println(arr1)
	fmt.Println("appending data")
	arr1 = append(arr1, 11)
	fmt.Println(arr1)
	var arr3 []int = []int{12, 13, 14, 15}
	fmt.Println(arr3)
	fmt.Println(arr1)
	arr1 = append(arr1, arr3...)
	fmt.Println(arr1)
	arr1 = append(arr1, 16, 17, 18, 19, 20)
	fmt.Println(arr1)
	fmt.Println("________________________________________________________________________")
	//Question 5
	arr4 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(arr4)
	arr5 := append(arr4[:3], arr4[6:]...)
	fmt.Println(arr5)
	fmt.Println("________________________________________________________________________")
	//Question 6
	arr6 := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(arr6)
	for i := 0; i < len(arr6); i++ {
		fmt.Println(i, ")", arr6[i], "length of the word is = ", len(arr6[i]))
	}
	fmt.Println("________________________________________________________________________")
	//Question 7
	arr7 := []string{"James", "Bond", "Shaken, not stirred"}
	arr8 := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	arr9 := [][]string{arr7, arr8}
	fmt.Println(arr7)
	fmt.Println(arr8)
	fmt.Println(arr9)
	for _, v1 := range arr9 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	fmt.Println("________________________________________________________________________")
	//Question 8
	map1 := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	if v, ok := map1[`bond_james`]; ok {
		fmt.Println(v)
	}
	for k, v := range map1 {

		fmt.Println(k)
		for _, v1 := range v {
			fmt.Printf("\t  %s", v1)
		}
		fmt.Printf("\n")
	}
	fmt.Println("________________________________________________________________________")
	//Question 9
	fmt.Println(map1)
	map1[`added_element`] = []string{`add`, `sub`, `mul`}
	fmt.Println(map1)
	fmt.Println("________________________________________________________________________")
	//Question 10
	delete(map1, `added_element`)
	fmt.Println(map1)
	fmt.Println(len(map1))
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
func otherBasic() {
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
