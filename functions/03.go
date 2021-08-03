package main 

import (
	"fmt"
)

func main() {
	//defer é um statment que sempre vai deixar aquela execucao por ultimo
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
}