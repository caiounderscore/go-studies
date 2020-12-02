package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 10; x++ {//inicializacao, condicacao, pós
		fmt.Println(string(x))
	}

	y := 0
	for { //loop infinito ate que seja dado o break
		if y < 10 {
			y++
			fmt.Println("y menor que 10, to dentro do loop")
		} else {
			fmt.Println("y maior que 10, vazei do loop")
			break
		}
	}

  for i := 33; i < 122; i++ {
    fmt.Printf("%#U - %v", i, i)
  }
}
