package main

import (
	"fmt"
)

func main() {
	x := "Pessoa"
	func(x string){ //outro exemplo de função anonima
		fmt.Println("olá", x)
	}(x)

	y := retornaumafunc() //adicionando uma função em uma variavel
	y()

	recebeumafunc(somefunction) //callback
}

func retornaumafunc() func() {
	return func(){
		fmt.Println("uepa")
	}
}

func somefunction() {
	fmt.Println("Colé Picolé")
}

func recebeumafunc(x func()) {
	fmt.Println("Retorno da func:")
	x()
}