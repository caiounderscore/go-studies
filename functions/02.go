package main

import (
	"fmt"
)

func main() {
	si := []int{10, 20, 30, 1, 2, 3}

	total := soma(si...) //Desenrolando uma slice.

	fmt.Println(total)
}

func soma(x ...int) int { // "..." deve ser obrigatoriamente o ultimo argumento
	var total int
	for _, v := range x {
		total += v
	}

	return total
}
