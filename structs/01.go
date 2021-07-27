package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	pessoa1 := pessoa{
		nome:      "Jõao",
		sobrenome: "das Couves",
		sabores:   []string{"chocolate", "morango"},
	}

	pessoa2 := pessoa{"Mario", "Bros", []string{"baunilha", "uva"}}

	fmt.Println(pessoa1, pessoa2)
	for _, v := range pessoa2.sabores {
		fmt.Println("\t", v)
	}
}
