package main

import (
	"fmt"
)

func main() {
	total, quantos := soma(10,10)
	fmt.Println(total, quantos)
}

// parametros  //retorno

//x, y,
func soma(x ...int) (int, int) {
	var total int
	for _, v := range x {
		total += v
	}
	return total, len(x)

}
