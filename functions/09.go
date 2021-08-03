package main

import ( 
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (p pessoa) diz() {
	fmt.Println("Olá meu nome é", p.nome, "meu sobrenome é", p.sobrenome, "e minha idade é", p.idade)
} 

func main() {
	p1 := pessoa{"Caio", "Rodrigues", 25}
	p1.diz()
}