package main

import "fmt"

const PRODUCT string = "Dr Congo"
const PRICE float64 = 500.2
const QUANTITY float64 = 27.982

var total float64

func main() {
	fmt.Println("The product is from", PRODUCT)
	fmt.Println("The price is :", PRICE)
	fmt.Println("The quantity is :", QUANTITY)
	fmt.Println("The total is :", PRICE*QUANTITY)

}
