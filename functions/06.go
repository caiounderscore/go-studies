package main

import (
	"fmt"
)

func main() {
	x := returnaumafuncao()
	y := x(3)
	fmt.Println(y)
	fmt.Println(returnaumafuncao()(3))
}

func returnaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}