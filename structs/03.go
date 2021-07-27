package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor	   string
}

type caminhonete struct {
	veiculo         
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo    
	modeloLuxo bool
}

func main() {
	ranger := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "azul",
		},
		tracaoNasQuatro: true,
	}

	siena := sedan{
		veiculo: veiculo{
			portas: 2,
			cor:    "prata",
		},
		modeloLuxo: false,
	}

	fmt.Println(ranger, siena)
	fmt.Println(siena.modeloLuxo,"\n",ranger.tracaoNasQuatro)
}
