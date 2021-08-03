package main 

import (
	"fmt"
)

func main() {
	a := contador()
	b := contador()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	//closure
}

func contador() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}