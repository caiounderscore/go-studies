package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	idade int 
}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

func main() {
	p1 := pessoa{
		nome: "Alice", 
		idade: 22,
	}

	p1.oibomdia()
}