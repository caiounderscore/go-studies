package main

import (
	"fmt"
)

func main() {
	qtoMundinho, mundinhoDeCaio := mundinho("caio")
	fmt.Println(mundinhoDeCaio)
	fmt.Println(qtoMundinho)

	fmt.Println(isPair(3))
}

func isPair(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

func mundinho(s string) (int, string) {
	return 1, "Olá mundinho de " + s
}