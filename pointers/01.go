package main

import (
	"fmt"
)

func main() {
	x := 10
	y := &x //endereço de memoria dessa variavel x

	fmt.Println(*y) //de-reference, mostra-se o que esta dentro do endereço
	fmt.Printf("%T, %T \n", x, y)
	*y++
	fmt.Println(x)
}