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
	meumapa := map[string]pessoa{
		"Das Couves": pessoa{
		nome:      "Jõao",
		sobrenome: "das Couves",
		sabores:   []string{"chocolate", "morango"},
		},
		"Bros": pessoa{"Mario", "Bros", []string{"baunilha", "uva"}},
	}

	for _, v := range meumapa {
		fmt.Println("Meu nome é", v.nome, "e meus sabores de sorvete favoritos são:")
		for _, v := range v.sabores {
			fmt.Println(v)
		}
		fmt.Println("\n")
	}
}
