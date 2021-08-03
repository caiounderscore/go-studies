package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 2
	func(a int, b int) { //funcao anonima, chama apenas uma vez
		fmt.Println(a * b)
	}(x, y)
}