package main

import "fmt"

type Currency struct {
	Name string 
	Country string 
	Number int 
}
var CAD = Currency{
	Name: "Canadian Dollar",
	Country: "Canada",
	Number : 124}
var USD = Currency{
	Name: "US Dollar", 
	Country: "USA", 
	Number: 840}

func main() {
	num0 := 242
	if num0 > 100 || num0 <  900 {
		fmt.Println("Currency:",num0)
		printCurr(num0)
	} else {
		fmt.Println("Currency unknown")
	}
	if num1 := 388; num1 > 100 || num1 < 900 {
		fmt.Println("Currency:",num1)
		printCurr(num1)
	}
}

func printCurr(number int) {
	if CAD.Number == number {
		fmt.Println("Found: %+v\n",CAD)
	} else {
		fmt.Println("No way!!!")
	}
}
