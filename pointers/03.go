package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {

	p1 := pessoa{
		"Caio",
		"Rodrigues",
		25,
	}

	fmt.Println(p1)

	mudeMe(&p1)

	fmt.Println(p1)
}

func mudeMe(p *pessoa) {
	(*p).nome = "UNDEFINIED"
}
