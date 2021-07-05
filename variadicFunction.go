package main 

import "fmt"

func main() {
	variadicExample("apple","mango","orange","banana")
}

func variadicExample(d ...string){
	fmt.Println(d[1])
	fmt.Println(d[3])
}
